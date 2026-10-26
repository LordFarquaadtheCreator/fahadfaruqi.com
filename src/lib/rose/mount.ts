import { PETALS, hash, type Rose, vnoise } from './build';
import RoseWorker from './rose.worker?worker&inline';

const DPR_CAP = 2;

// The camera. The flower is ~3.1 units across; FIT maps that onto 80% of the shorter edge of the
// window, and the distance gives enough perspective that turning it reads as depth.
const CAM_DISTANCE = 5.2;
const FIT = 1.34;
const CENTER_Y = 0.75;
const YAW_PER_PX = 0.008;
const PITCH_PER_PX = 0.008;
const PITCH_MIN = -0.45;
const PITCH_MAX = 1.35;

// One occlusion cell per this many CSS pixels: enough to hide the far side of the bloom without a
// full-resolution depth buffer.
const OCCLUSION_CELL = 6;
const DEPTH_EPSILON = 0.035;

const LIGHT = { x: -0.38, y: 0.3, z: 0.87 };
const LIGHT_LEN = Math.sqrt(LIGHT.x * LIGHT.x + LIGHT.y * LIGHT.y + LIGHT.z * LIGHT.z);
LIGHT.x /= LIGHT_LEN;
LIGHT.y /= LIGHT_LEN;
LIGHT.z /= LIGHT_LEN;

function clamp(x: number) {
	return Math.min(1, Math.max(0, x));
}

/**
 * Draws the rose for one seed. The geometry is built in a worker; the view is rotary — drag the
 * canvas to turn it, and every mark is re-lit and re-classified from its rotated normal, so the
 * etching follows the surface instead of a fixed frame.
 */
export function mountRose(
	canvas: HTMLCanvasElement,
	seed: number,
	onDrawn: () => void = () => {}
): () => void {
	const got = canvas.getContext('2d');
	if (!got) return () => {};
	const ctx = got;

	let disposed = false;
	let raf = 0;
	let rose: Rose | null = null;

	// the view
	let yaw = 0;
	let pitch = 0.62;
	let dragging = false;
	let lastX = 0;
	let lastY = 0;

	// derived from the canvas size
	let focal = 1;
	let cx = 0;
	let cy = 0;
	let cellPx = OCCLUSION_CELL;
	let gridW = 1;
	let gridH = 1;
	let depthGrid = new Float32Array(1);
	let unit = 6;

	// per-mark scratch, allocated once per drawing
	let pxBuf = new Float32Array(1);
	let pyBuf = new Float32Array(1);
	let depthBuf = new Float32Array(1);
	let nxBuf = new Float32Array(1);
	let nyBuf = new Float32Array(1);
	let nzBuf = new Float32Array(1);

	const quiet = matchMedia('(prefers-reduced-motion: reduce)').matches;

	function layout(): boolean {
		const cssW = canvas.clientWidth;
		const cssH = canvas.clientHeight;
		if (cssW < 1 || cssH < 1) return false;
		const dpr = Math.min(window.devicePixelRatio || 1, DPR_CAP);
		const w = Math.round(cssW * dpr);
		const h = Math.round(cssH * dpr);
		if (canvas.width !== w || canvas.height !== h) {
			canvas.width = w;
			canvas.height = h;
		}
		const minDim = Math.min(w, h);
		focal = FIT * minDim;
		cx = w / 2;
		cy = h / 2;
		cellPx = Math.max(2, OCCLUSION_CELL * dpr);
		gridW = Math.ceil(w / cellPx) + 1;
		gridH = Math.ceil(h / cellPx) + 1;
		if (depthGrid.length !== gridW * gridH) depthGrid = new Float32Array(gridW * gridH);
		// one unit of the sketch's six-pixel cell, so mark lengths keep their proportions
		unit = (minDim / 900) * 6;
		return true;
	}

	function ensureScratch(count: number) {
		if (pxBuf.length === count) return;
		pxBuf = new Float32Array(count);
		pyBuf = new Float32Array(count);
		depthBuf = new Float32Array(count);
		nxBuf = new Float32Array(count);
		nyBuf = new Float32Array(count);
		nzBuf = new Float32Array(count);
	}

	function draw(ms: number) {
		const data = rose;
		if (!data || !layout()) return;
		ensureScratch(data.count);

		// the wind: one shared period, outer petals leading, each petal swinging about the stem
		const t = ms / 1000;
		const AMP = 0.05;
		const cosK = new Float32Array(PETALS);
		const sinK = new Float32Array(PETALS);
		for (let k = 0; k < PETALS; k++) {
			const r = hash(k, 23),
				T = 9,
				ph = ((((t - 2.4 * (1 - k / (PETALS - 1)) - 0.1 * r) % T) + T) % T) / T,
				wv = 0.5 - 0.5 * Math.cos(2 * Math.PI * ph),
				e = wv * wv * (3 - 2 * wv),
				breathe =
					e * (0.75 + 0.5 * vnoise(k * 0.03 + 5, t * 0.25)) * (1 + 0.12 * Math.sin(t * 1.1 + k * 0.15)),
				ang = breathe * AMP;
			cosK[k] = Math.cos(ang);
			sinK[k] = Math.sin(ang);
		}

		const cyaw = Math.cos(yaw),
			syaw = Math.sin(yaw),
			cp = Math.cos(pitch),
			sp = Math.sin(pitch);

		const n = data.count;
		const rx = data.x,
			ry = data.y,
			rz = data.z,
			rnx = data.nx,
			rny = data.ny,
			rnz = data.nz,
			petal = data.petal;

		depthGrid.fill(-1e9);

		// pass 1: place every mark, keep the nearest one per occlusion cell
		for (let i = 0; i < n; i++) {
			const k = petal[i],
				ck = cosK[k],
				sk = sinK[k];
			// the wind, about the stem
			const wx = rx[i] * ck - rz[i] * sk,
				wz = rx[i] * sk + rz[i] * ck,
				wy = ry[i];
			const wnx = rnx[i] * ck - rnz[i] * sk,
				wnz = rnx[i] * sk + rnz[i] * ck,
				wny = rny[i];
			// camera: yaw, then pitch
			const x1 = wx * cyaw - wz * syaw;
			const z1 = wx * syaw + wz * cyaw;
			const y2 = wy * cp - z1 * sp;
			const z2 = wy * sp + z1 * cp;
			const ux1 = wnx * cyaw - wnz * syaw;
			const uz1 = wnx * syaw + wnz * cyaw;
			const uy2 = wny * cp - uz1 * sp;
			const uz2 = wny * sp + uz1 * cp;

			const depth = CAM_DISTANCE - z2;
			const scale = focal / depth;
			pxBuf[i] = cx + x1 * scale;
			pyBuf[i] = cy - (y2 - CENTER_Y) * scale;
			depthBuf[i] = depth;
			nxBuf[i] = ux1;
			nyBuf[i] = uy2;
			nzBuf[i] = uz2;

			const gx = (pxBuf[i] / cellPx) | 0,
				gy = (pyBuf[i] / cellPx) | 0;
			if (gx >= 0 && gy >= 0 && gx < gridW && gy < gridH) {
				const gi = gy * gridW + gx;
				const closeness = -depth;
				if (depthGrid[gi] < closeness) depthGrid[gi] = closeness;
			}
		}

		// pass 2: only the marks this view can actually see, dressed by their own normal
		const paths: Path2D[] = [];
		for (let b = 0; b <= 16; b++) paths.push(new Path2D());
		for (let i = 0; i < n; i++) {
			const gx = (pxBuf[i] / cellPx) | 0,
				gy = (pyBuf[i] / cellPx) | 0;
			if (gx < 0 || gy < 0 || gx >= gridW || gy >= gridH) continue;
			const depth = depthBuf[i];
			if (-depth < depthGrid[gy * gridW + gx] - DEPTH_EPSILON) continue;

			const vnx = nxBuf[i],
				vny = nyBuf[i],
				vnz = nzBuf[i];
			const lam = Math.pow(clamp((vnx * LIGHT.x + vny * LIGHT.y + vnz * LIGHT.z + 0.45) / 1.45), 1.2);
			const fill = Math.max(0, vnx * 0.5 + vny * 0.2 + vnz * 0.84); // weak fill from the right
			const facing = Math.abs(vnz);
			const rim = Math.pow(1 - facing, 3) * 0.2;
			const L = clamp(0.12 + 0.6 * lam + 0.1 * fill + rim);

			// a mark near the camera reads heavier than the same mark on the far side
			const u = unit * (CAM_DISTANCE / depth);
			const x = pxBuf[i],
				y = pyBuf[i];

			// the rim of the surface: crosses, whatever the light is doing
			if (facing < 0.22) {
				const arm = 2.6 * u;
				const p = paths[Math.max(1, Math.round((0.5 + 0.5 * Math.max(L, 0.35)) * 15))];
				p.moveTo(x - arm, y);
				p.lineTo(x + arm, y);
				p.moveTo(x, y - arm);
				p.lineTo(x, y + arm);
				continue;
			}

			let len: number, bucket: number;
			if (L > 0.6) {
				len = 3.4 * u; // highlight: the long dash, and the bold register
				bucket = 16;
			} else if (L > 0.5) {
				len = (1.2 + 2.2 * L) * u;
				bucket = Math.max(1, Math.round((0.55 + 0.45 * L) * 15));
			} else if (L > 0.38) {
				len = (0.9 + 1.4 * L) * u;
				bucket = Math.max(1, Math.round((0.3 + 0.7 * L) * 15));
			} else if (L > 0.22) {
				len = 0.8 * u;
				bucket = Math.max(1, Math.round((0.25 + 0.8 * L) * 15));
			} else {
				len = 0.7 * u; // shadow: the shortest mark there is
				bucket = Math.max(1, Math.round((0.2 + L) * 15));
			}
			const p = paths[bucket];
			p.moveTo(x - len, y);
			p.lineTo(x + len, y);
		}

		const w = canvas.width;
		const h = canvas.height;
		ctx.setTransform(1, 0, 0, 1, 0, 0);
		ctx.globalCompositeOperation = 'source-over';
		ctx.globalAlpha = 1;
		ctx.clearRect(0, 0, w, h);
		// The marks are added onto the window's own ground. The bloom is a second, wider and fainter
		// pass of the same paths: a halo costs no extra pixels to composite, where downscaling the
		// art three times and stretching it back cost nine full-canvas blits a frame.
		ctx.globalCompositeOperation = 'lighter';
		ctx.lineWidth = Math.max(2, unit * 0.9);
		for (let b = 1; b <= 15; b++) {
			ctx.strokeStyle = `rgba(255,255,255,${((b / 15) * 0.22).toFixed(3)})`;
			ctx.stroke(paths[b]);
		}
		ctx.lineWidth = Math.max(1, unit * 0.22);
		for (let b = 1; b <= 15; b++) {
			ctx.strokeStyle = `rgba(255,255,255,${(b / 15).toFixed(3)})`;
			ctx.stroke(paths[b]);
		}
		ctx.lineWidth = Math.max(1.4, unit * 0.32);
		ctx.strokeStyle = '#fff';
		ctx.stroke(paths[16]);
	}

	function tick(ms: number) {
		if (disposed) return;
		draw(ms);
		raf = requestAnimationFrame(tick);
	}

	function onPointerDown(event: PointerEvent) {
		dragging = true;
		lastX = event.clientX;
		lastY = event.clientY;
		canvas.setPointerCapture(event.pointerId);
		canvas.classList.add('rose-canvas--dragging');
	}

	function onPointerMove(event: PointerEvent) {
		if (!dragging) return;
		yaw += (event.clientX - lastX) * YAW_PER_PX;
		pitch = Math.min(PITCH_MAX, Math.max(PITCH_MIN, pitch + (event.clientY - lastY) * PITCH_PER_PX));
		lastX = event.clientX;
		lastY = event.clientY;
		if (quiet) draw(performance.now());
	}

	function onPointerUp(event: PointerEvent) {
		dragging = false;
		if (canvas.hasPointerCapture(event.pointerId)) canvas.releasePointerCapture(event.pointerId);
		canvas.classList.remove('rose-canvas--dragging');
	}

	const observer = new ResizeObserver(() => {
		if (!rose || quiet) draw(0);
	});
	observer.observe(canvas);

	canvas.addEventListener('pointerdown', onPointerDown);
	canvas.addEventListener('pointermove', onPointerMove);
	canvas.addEventListener('pointerup', onPointerUp);
	canvas.addEventListener('pointercancel', onPointerUp);

	const worker = new RoseWorker();
	worker.onmessage = (event: MessageEvent<Rose>) => {
		if (disposed) return;
		rose = event.data;
		ensureScratch(rose.count);
		worker.terminate();
		if (quiet) draw(0);
		else raf = requestAnimationFrame(tick);
		onDrawn();
	};
	worker.postMessage({ seed });

	return () => {
		disposed = true;
		cancelAnimationFrame(raf);
		canvas.removeEventListener('pointerdown', onPointerDown);
		canvas.removeEventListener('pointermove', onPointerMove);
		canvas.removeEventListener('pointerup', onPointerUp);
		canvas.removeEventListener('pointercancel', onPointerUp);
		observer.disconnect();
		worker.terminate();
	};
}
