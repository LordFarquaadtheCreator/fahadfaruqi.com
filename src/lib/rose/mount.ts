import { ROSE_H, ROSE_SCALE, ROSE_W, type Rose } from './build';
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

const YAW_PER_PX = 0.008;
const PITCH_PER_PX = 0.006;
const PITCH_MIN = -0.55;
const PITCH_MAX = 1.2;

// Occlusion is decided on a coarse grid in screen pixels: enough to keep the far side of the bloom
// behind the near one.
const OCCLUSION_CELL = 5;
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
	let rose: Rose | null = null;

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
	let unit = 6;

	let pxBuf = new Float32Array(1);
	let pyBuf = new Float32Array(1);
	let depthBuf = new Float32Array(1);

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
		if (depthGrid.length !== gridW * gridH) depthGrid = new Float32Array(gridW * gridH);
		unit = k; // mark lengths are built in art pixels
		return true;
	}

	function ensureScratch(count: number) {
		if (pxBuf.length === count) return;
		pxBuf = new Float32Array(count);
		pyBuf = new Float32Array(count);
		depthBuf = new Float32Array(count);
	}

	function draw() {
		const data = rose;
		if (!data || !layout()) return;

		const cy = Math.cos(yaw),
			sy = Math.sin(yaw);
		// the camera's tilt, plus whatever the drag has added
		const A = AL + pitch;
		const cA = Math.cos(A),
			sA = Math.sin(A);

		const marks = data.marks;
		const n = marks.length;
		depthGrid.fill(-1e9);

		// pass 1: turn, project, and keep the nearest depth per occlusion cell
		for (let i = 0; i < n; i++) {
			const m = marks[i];
			// spin about the flower's own axis
			const x1 = m.x * cy + m.z * sy;
			const z1 = -m.x * sy + m.z * cy;
			// the sketch's tilt, then its lean
			const y2 = m.y * cA - z1 * sA;
			const z2 = z1 * cA + m.y * sA;
			const cx3 = x1 * CG + y2 * SG;
			const cy3 = -x1 * SG + y2 * CG;

			pxBuf[i] = ox + (BX + cx3 * ROSE_SCALE) * k;
			pyBuf[i] = oy + (BY - cy3 * ROSE_SCALE) * k;
			depthBuf[i] = z2;

			const gx = (pxBuf[i] / cellPx) | 0,
				gy = (pyBuf[i] / cellPx) | 0;
			if (gx >= 0 && gy >= 0 && gx < gridW && gy < gridH) {
				const gi = gy * gridW + gx;
				if (depthGrid[gi] < z2) depthGrid[gi] = z2;
			}
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
			const path = paths[m.b];
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
		draw();
	}

	function onPointerUp(event: PointerEvent) {
		dragging = false;
		if (canvas.hasPointerCapture(event.pointerId)) canvas.releasePointerCapture(event.pointerId);
		canvas.classList.remove('rose-canvas--dragging');
	}

	const observer = new ResizeObserver(() => draw());
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
		draw();
		onDrawn();
	};
	worker.postMessage({ seed });

	return () => {
		disposed = true;
		canvas.removeEventListener('pointerdown', onPointerDown);
		canvas.removeEventListener('pointermove', onPointerMove);
		canvas.removeEventListener('pointerup', onPointerUp);
		canvas.removeEventListener('pointercancel', onPointerUp);
		observer.disconnect();
		worker.terminate();
	};
}
