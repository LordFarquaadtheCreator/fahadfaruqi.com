import { PETALS, ROSE_H, ROSE_SCALE, ROSE_W, hash, type Rose, vnoise } from './build';
import RoseWorker from './rose.worker?worker&inline';

const DPR_CAP = 2;
// The sketch's own camera: it looks down into the bloom by AL and leans by GM. The drag adds yaw
// about the flower's axis and pitch on top of AL; nothing else about the projection changes, so at
// rest the marks land exactly where the drawing puts them.
const AL = 0.8;
const GM = 0.16;
const CA = Math.cos(AL);
const SA = Math.sin(AL);
const CG = Math.cos(GM);
const SG = Math.sin(GM);
const BX = 470;
const BY = 700;

// The light stands still in the room, upper left and in front, exactly as the sketch lit it. The
// flower turns through it.
let LX = -0.38,
	LY = 0.3,
	LZ = 0.87;
const LL = Math.sqrt(LX * LX + LY * LY + LZ * LZ);
LX /= LL;
LY /= LL;
LZ /= LL;

function clamp(x: number) {
	return Math.min(1, Math.max(0, x));
}

// A 256-entry table for x^1.2, because this runs twelve thousand times a frame and Math.pow does
// not — the bucket it feeds is an integer, so interpolating would be wasted precision.
const POW12 = new Float32Array(257);
for (let i = 0; i <= 256; i++) POW12[i] = Math.pow(i / 256, 1.2);

/** The sketch's light formula, evaluated on a camera-space normal. */
function lightOf(nx: number, ny: number, nz: number, lit: number, noise: number) {
	const lam = POW12[(clamp((nx * LX + ny * LY + nz * LZ + 0.45) / 1.45) * 256) | 0] * noise;
	const fill = Math.max(0, nx * 0.5 + ny * 0.2 + nz * 0.84);
	const k = 1 - Math.max(0, nz);
	const rim = k * k * k * 0.2;
	return clamp(0.08 + 0.46 * lam * (0.4 + 0.6 * lit) + 0.1 * fill + rim);
}

/** A mark's normal turned into camera space for a given yaw and tilt: the lean does not change
 *  how much light the normal takes, because the lamp's z axis and the lean's axis coincide. */
function cameraNormal(nx: number, ny: number, nz: number, cy: number, sy: number, cA: number, sA: number, out: number[]) {
	const x1 = nx * cy + nz * sy;
	const z1 = -nx * sy + nz * cy;
	const y2 = ny * cA - z1 * sA;
	const z2 = z1 * cA + ny * sA;
	const cx = x1 * CG + y2 * SG;
	const cyy = -x1 * SG + y2 * CG;
	out[0] = cx;
	out[1] = cyy;
	out[2] = z2;
}

const YAW_PER_PX = 0.008;
const PITCH_PER_PX = 0.006;
const PITCH_MIN = -0.55;
const PITCH_MAX = 1.2;

// Occlusion is decided on a fine grid in screen pixels, then grown by one cell: marks are sparse on
// screen, so without that growth the back of the flower shows through the gaps between the front
// ones. Two CSS pixels, because a coarser cell cannot separate a petal from the one behind it.
const OCCLUSION_CELL = 2;
const DEPTH_EPSILON = 0.02;

/**
 * Draws the rose for one seed. The drawing — every mark, its length and its weight — is built in a
 * worker exactly as the sketch builds it; this turns those marks in space and projects them, so the
 * resting view is the drawing itself and dragging is a rotation of it.
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
	// read here, not at module scope: the module is also evaluated during the prerender, where
	// there is no matchMedia
	const quiet = matchMedia('(prefers-reduced-motion: reduce)').matches;

	// the view
	let yaw = 0;
	let pitch = 0;
	let dragging = false;
	let lastX = 0;
	let lastY = 0;

	// derived from the canvas size
	let k = 1; // art pixels → canvas pixels
	let ox = 0;
	let oy = 0;
	let cellPx = OCCLUSION_CELL;
	let gridW = 1;
	let gridH = 1;
	let depthGrid = new Float32Array(1);
	let depthGrown = new Float32Array(1);
	let unit = 6;

	let pxBuf = new Float32Array(1);
	let pyBuf = new Float32Array(1);
	let depthBuf = new Float32Array(1);
	let bucketBuf = new Uint8Array(1);
	let restLight = new Float32Array(1);

	// The drawing no longer animates: it is redrawn when the view changes, not sixty times a second.

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
		// the drawing is fitted into the window the way the sketch's own frame was
		k = Math.min(w / ROSE_W, h / ROSE_H);
		ox = (w - ROSE_W * k) / 2;
		oy = (h - ROSE_H * k) / 2;
		cellPx = Math.max(2, OCCLUSION_CELL * dpr);
		gridW = Math.ceil(w / cellPx) + 1;
		gridH = Math.ceil(h / cellPx) + 1;
		if (depthGrid.length !== gridW * gridH) {
			depthGrid = new Float32Array(gridW * gridH);
			depthGrown = new Float32Array(gridW * gridH);
		}
		unit = k; // mark lengths are built in art pixels
		return true;
	}

	function ensureScratch(count: number) {
		if (pxBuf.length === count) return;
		pxBuf = new Float32Array(count);
		pyBuf = new Float32Array(count);
		depthBuf = new Float32Array(count);
		bucketBuf = new Uint8Array(count);
		// what the lamp gave each mark at rest: the drawing's own shading, in camera space at last
		restLight = new Float32Array(count);
		const sA0 = Math.sin(AL),
			cA0 = Math.cos(AL);
		const n: number[] = [0, 0, 0];
		for (let i = 0; i < count; i++) {
			const m = rose!.marks[i];
			cameraNormal(m.nx, m.ny, m.nz, 1, 0, cA0, sA0, n);
			restLight[i] = Math.max(0.02, lightOf(n[0], n[1], n[2], m.lit, m.noise));
		}
	}

	function draw(now: number) {
		const data = rose;
		if (!data || !layout()) return;

		const t = now / 1000;
		// the sketch's opening cycle: one shared nine-second period, outer petals leading, each
		// petal swinging out from the base and coming back — in 3D, about the axis across its radius
		const AMAX = 0.11;
		const cosK = new Float32Array(PETALS);
		const sinK = new Float32Array(PETALS);
		for (let k = 0; k < PETALS; k++) {
			const r = hash(k, 23),
				T = 9,
				ph = ((((t - 2.4 * (1 - k / (PETALS - 1)) - 0.1 * r) % T) + T) % T) / T,
				wv = 0.5 - 0.5 * Math.cos(2 * Math.PI * ph),
				e = wv * wv * (3 - 2 * wv);
			const breathe =
				e * (0.75 + 0.5 * vnoise(k * 0.03 + 5, t * 0.25)) * (1 + 0.12 * Math.sin(t * 1.1 + k * 0.15));
			const ang = breathe * AMAX;
			cosK[k] = Math.cos(ang);
			sinK[k] = Math.sin(ang);
		}

		const cy = Math.cos(yaw),
			sy = Math.sin(yaw);
		// the camera's tilt, plus whatever the drag has added
		const A = AL + pitch;
		const cA = Math.cos(A),
			sA = Math.sin(A);

		const marks = data.marks;
		const n = marks.length;
		depthGrid.fill(-1e9);

		// pass 1: swing, turn, project, light; keep the nearest depth per occlusion cell
		for (let i = 0; i < n; i++) {
			const m = marks[i];
			// the petal's swing, about a horizontal axis across its radius
			let px = m.x,
				py = m.y,
				pz = m.z;
			let nx = m.nx,
				ny = m.ny,
				nz = m.nz;
			if (m.spread > 0.001) {
				// the angle is per petal, so recover it from the mark's own axis: every petal shares
				// one axis, and the spread scales that petal's angle
				const pk = m.petal;
				const c = 1 + (cosK[pk] - 1) * m.spread;
				const s = sinK[pk] * m.spread;
				// Rodrigues about (ax, ay, az)
				const d = m.ax * px + m.ay * py + m.az * pz;
				const cxp = m.ay * pz - m.az * py,
					cyp = m.az * px - m.ax * pz,
					czp = m.ax * py - m.ay * px;
				px = px * c + cxp * s + m.ax * d * (1 - c);
				py = py * c + cyp * s + m.ay * d * (1 - c);
				pz = pz * c + czp * s + m.az * d * (1 - c);
				const dn = m.ax * nx + m.ay * ny + m.az * nz;
				const cxn = m.ay * nz - m.az * ny,
					cyn = m.az * nx - m.ax * nz,
					czn = m.ax * ny - m.ay * nx;
				nx = nx * c + cxn * s + m.ax * dn * (1 - c);
				ny = ny * c + cyn * s + m.ay * dn * (1 - c);
				nz = nz * c + czn * s + m.az * dn * (1 - c);
			}
			// spin about the flower's own axis, then the sketch's camera
			const x1 = px * cy + pz * sy;
			const z1 = -px * sy + pz * cy;
			const n1x = nx * cy + nz * sy;
			const n1z = -nx * sy + nz * cy;

			const y2 = py * cA - z1 * sA;
			const z2 = z1 * cA + py * sA;
			const nx2 = n1x;
			const ny2 = ny * cA - n1z * sA;
			const nz2 = n1z * cA + ny * sA;

			const cx3 = x1 * CG + y2 * SG;
			const cy3 = -x1 * SG + y2 * CG;
			const ncx = nx2 * CG + ny2 * SG;
			const ncy = -nx2 * SG + ny2 * CG;

			pxBuf[i] = ox + (BX + cx3 * ROSE_SCALE) * k;
			pyBuf[i] = oy + (BY - cy3 * ROSE_SCALE) * k;
			depthBuf[i] = z2;

			// The lamp stands still in the room; the flower turns through it, so each mark's ink is
			// scaled by the light it catches now against the light it caught at rest. At rest the
			// ratio is one and the drawing is exactly the drawing; turned, the marks that face the
			// lamp carry the ink and the ones that leave it fade out.
			const ratio = lightOf(ncx, ncy, nz2, m.lit, m.noise) / restLight[i];
			bucketBuf[i] = Math.max(0, Math.min(16, Math.round(m.b * Math.pow(ratio, 1.5))));

			const gx = (pxBuf[i] / cellPx) | 0,
				gy = (pyBuf[i] / cellPx) | 0;
			if (gx >= 0 && gy >= 0 && gx < gridW && gy < gridH) {
				const gi = gy * gridW + gx;
				if (depthGrid[gi] < z2) depthGrid[gi] = z2;
			}
		}

		// the nearest depth within one cell of each cell: marks sit in the gaps between other marks,
		// so a one-cell lookup would let the far side of the bloom through
		// separable: the three-wide maximum along the row, then down the column
		for (let gy = 0; gy < gridH; gy++) {
			const row = gy * gridW;
			for (let gx = 0; gx < gridW; gx++) {
				let best = depthGrid[row + gx];
				if (gx > 0 && depthGrid[row + gx - 1] > best) best = depthGrid[row + gx - 1];
				if (gx < gridW - 1 && depthGrid[row + gx + 1] > best) best = depthGrid[row + gx + 1];
				depthGrown[row + gx] = best;
			}
		}
		for (let gx = 0; gx < gridW; gx++)
			for (let gy = 0; gy < gridH; gy++) {
				const i = gy * gridW + gx;
				let best = depthGrown[i];
				if (gy > 0 && depthGrown[i - gridW] > best) best = depthGrown[i - gridW];
				if (gy < gridH - 1 && depthGrown[i + gridW] > best) best = depthGrown[i + gridW];
				depthGrid[i] = best;
			}

		// pass 2: the marks this view can see, drawn as the drawing draws them
		const paths: Path2D[] = [];
		for (let b = 0; b <= 16; b++) paths.push(new Path2D());
		for (let i = 0; i < n; i++) {
			const gx = (pxBuf[i] / cellPx) | 0,
				gy = (pyBuf[i] / cellPx) | 0;
			if (gx < 0 || gy < 0 || gx >= gridW || gy >= gridH) continue;
			if (depthBuf[i] < depthGrid[gy * gridW + gx] - DEPTH_EPSILON) continue;

			const m = marks[i];
			const x = pxBuf[i],
				y = pyBuf[i];
			const len = m.n * ROSE_SCALE * unit;
			const bkt = bucketBuf[i];
			if (bkt < 1) continue;
			const path = paths[bkt];
			if (m.t === 0 || m.t === 3) {
				path.moveTo(x - len, y);
				path.lineTo(x + len, y);
				path.moveTo(x, y - len);
				path.lineTo(x, y + len);
			} else {
				path.moveTo(x - len, y);
				path.lineTo(x + len, y);
			}
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
		ctx.lineWidth = Math.max(2, unit * 2.4);
		for (let b = 1; b <= 15; b++) {
			ctx.strokeStyle = `rgba(255,255,255,${((b / 15) * 0.22).toFixed(3)})`;
			ctx.stroke(paths[b]);
		}
		ctx.lineWidth = Math.max(1, unit * 1.3);
		for (let b = 1; b <= 15; b++) {
			ctx.strokeStyle = `rgba(255,255,255,${(b / 15).toFixed(3)})`;
			ctx.stroke(paths[b]);
		}
		ctx.lineWidth = Math.max(1.4, unit * 1.9);
		ctx.strokeStyle = '#fff';
		ctx.stroke(paths[16]);
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

	function tick(now: number) {
		if (disposed) return;
		draw(now);
		raf = requestAnimationFrame(tick);
	}

	function onPointerUp(event: PointerEvent) {
		dragging = false;
		if (canvas.hasPointerCapture(event.pointerId)) canvas.releasePointerCapture(event.pointerId);
		canvas.classList.remove('rose-canvas--dragging');
	}

	const observer = new ResizeObserver(() => {
		if (!rose || quiet) draw(performance.now());
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
