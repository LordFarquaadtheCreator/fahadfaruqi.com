// The rose, built exactly as the sketch builds it: a shadow map, a depth buffer, then one mark per
// six-pixel cell decided by the light, the silhouette and the petal lips. The one addition is that
// every mark remembers the 3D point it came from, so the renderer turns the drawing rather than
// redrawing it.
//
// The buffers are coarser than the sketch's own — three art pixels to a buffer pixel, a quarter of
// its surface samples — because detail that fine only ever decides which of two nearly identical
// marks a cell gets. The mark set is the same, and the build is twenty times cheaper.

export const ROSE_W = 900;
export const ROSE_H = 1040;
/** Camera units to art pixels. */
export const ROSE_SCALE = 380;
const BX = 470;
const BY = 700;
const CELL = 6;
/** Buffer pixels per art pixel. */
const STEP = 1;
const BW = Math.ceil(ROSE_W / STEP);
const BH = Math.ceil(ROSE_H / STEP);
/** Samples across each petal and along it. */
const NU = 220;
const NV = 250;
const STEM_LENGTH = 0.45;
export const PETALS = 44;
const N = PETALS;

// the sketch's camera: tilted down into the bloom, leaning a little
const AL = 0.8;
const GM = 0.16;
const CA = Math.cos(AL);
const SA = Math.sin(AL);
const CG = Math.cos(GM);
const SG = Math.sin(GM);

/** 0 cross, 1 dash, 2 dot, 3 bold cross. */
export type MarkKind = 0 | 1 | 2 | 3;

export interface RoseMark {
	t: MarkKind;
	/** Flower-local position: x right, y up the stem, z toward the viewer at yaw 0. */
	x: number;
	y: number;
	z: number;
	/** Flower-local surface normal: the renderer relights with it as the flower turns. */
	nx: number;
	ny: number;
	nz: number;
	/** Half the stroke's length, in camera units. */
	n: number;
	/** Alpha bucket, 1..16 (16 is the bold register). */
	b: number;
	/** How much of the mark the lamp could reach, from the shadow map: a property of the light, not
	 *  of the view, so it holds as the flower turns. */
	lit: number;
	/** The sketch's texture noise at this cell. */
	noise: number;
	/** The petal it belongs to, the axis that petal swings about, and how far out it sits. */
	petal: number;
	ax: number;
	ay: number;
	az: number;
	spread: number;
}

export interface Rose {
	count: number;
	marks: RoseMark[];
}

export function rng(seed: number) {
	let s = seed;
	return function () {
		s |= 0;
		s = (s + 0x6d2b79f5) | 0;
		let t = Math.imul(s ^ (s >>> 15), 1 | s);
		t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
		return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
	};
}

export function hash(i: number, j: number) {
	const s = Math.sin(i * 127.1 + j * 311.7) * 43758.5453;
	return s - Math.floor(s);
}

function lerp(p: number, q: number, t: number) {
	return p + (q - p) * t;
}

export function vnoise(x: number, y: number) {
	const i = Math.floor(x),
		j = Math.floor(y),
		fx = x - i,
		fy = y - j;
	const u = fx * fx * (3 - 2 * fx),
		v = fy * fy * (3 - 2 * fy);
	return lerp(lerp(hash(i, j), hash(i + 1, j), u), lerp(hash(i, j + 1), hash(i + 1, j + 1), u), v);
}

function clamp(x: number) {
	return Math.min(1, Math.max(0, x));
}

interface Petal {
	phi: number;
	w: number;
	r0: number;
	b0: number;
	b1: number;
	L: number;
	y0: number;
	tw: number;
	ph: number;
}

interface Point {
	x: number;
	y: number;
	z: number;
}

export function buildRose(seed: number): Rose {
	const R = rng(seed);
	const petals: Petal[] = [];
	for (let i = 0; i < N; i++) {
		const t = i / (N - 1),
			r0 = 0.02 + 0.08 * t,
			tt = Math.pow(t, 0.9);
		petals.push({
			phi: i * 2.39996 + R() * 0.2,
			w: 1.25 - 0.2 * t + R() * 0.15,
			r0,
			b0: 0.12 + 0.6 * tt + R() * 0.06,
			b1: -0.1 + 0.3 * tt + R() * 0.08,
			L: 1.5 - 0.35 * t + R() * 0.12,
			y0: 0.02 * t,
			tw: R() - 0.5,
			ph: R() * 6
		});
	}

	// one petal's surface in the flower's own frame
	function pos(p: Petal, u: number, v: number, o: Point) {
		const vv = Math.max(v, 0),
			w = p.w * (0.5 + 0.5 * Math.sqrt(vv)),
			phi = p.phi + u * w + 0.25 * p.tw * v;
		let k = p.b1 - p.b0;
		if (Math.abs(k) < 1e-3) k = 1e-3;
		const bb = p.b0 + k * vv;
		const rho =
			p.r0 +
			(p.L * (Math.cos(p.b0) - Math.cos(bb))) / k +
			0.02 * v * Math.sin(u * 4 + p.ph) +
			0.012 * vv * Math.sin(u * 13 + p.ph * 2);
		const y = p.y0 + (p.L * (Math.sin(bb) - Math.sin(p.b0))) / k - 0.08 * u * u * vv;
		o.x = rho * Math.sin(phi);
		o.y = y;
		o.z = rho * Math.cos(phi);
	}

	// the sketch's camera: tilt into the bloom, lean, then an orthographic projection
	function toCamera(p: Point, o: Point) {
		const y2 = p.y * CA - p.z * SA,
			z2 = p.z * CA + p.y * SA;
		o.x = p.x * CG + y2 * SG;
		o.y = -p.x * SG + y2 * CG;
		o.z = z2;
	}

	const D4x = [1, -1, 0, 0],
		D4y = [0, 0, 1, -1];
	const A: Point = { x: 0, y: 0, z: 0 };
	const B: Point = { x: 0, y: 0, z: 0 };
	const C: Point = { x: 0, y: 0, z: 0 };
	const CAM: Point = { x: 0, y: 0, z: 0 };
	const CAM2: Point = { x: 0, y: 0, z: 0 };

	// ---- 1. the light, and its own depth buffer
	let Lx = -0.38,
		Ly = 0.3,
		Lz = 0.87;
	const ll = Math.sqrt(Lx * Lx + Ly * Ly + Lz * Lz);
	Lx /= ll;
	Ly /= ll;
	Lz /= ll;
	let E1x = Lz,
		E1z = -Lx;
	const el = Math.sqrt(E1x * E1x + E1z * E1z);
	E1x /= el;
	E1z /= el;
	const E2x = Ly * E1z,
		E2y = Lz * E1x - Lx * E1z,
		E2z = -Ly * E1x;
	const SMW = 1600,
		SMC = 800,
		SMS = ROSE_SCALE,
		SM = new Float32Array(SMW * SMW).fill(-1e9);

	function eachSample(
		withNormal: boolean,
		cb: (idx: number, local: Point, cam: Point, petal: number, v: number, normal: Point | null) => void
	) {
		for (let k = 0; k < N; k++) {
			const p = petals[k];
			for (let iu = 0; iu <= NU; iu++) {
				const u = -1 + (2 * iu) / NU,
					vmax = 1 - 0.3 * u * u + 0.04 * Math.sin(u * 3 + p.ph);
				for (let iv = 0; iv <= NV; iv++) {
					const v = (iv / NV) * vmax;
					const idx = (k * (NU + 1) + iu) * (NV + 1) + iv;
					pos(p, u, v, A);
					toCamera(A, CAM);
					if (!withNormal) {
						cb(idx, A, CAM, k, v, null);
						continue;
					}
					pos(p, u + 0.01, v, B);
					pos(p, u, v + 0.01, C);
					const ax = B.x - A.x,
						ay = B.y - A.y,
						az = B.z - A.z,
						bx = C.x - A.x,
						by = C.y - A.y,
						bz = C.z - A.z;
					let nx = ay * bz - az * by,
						ny = az * bx - ax * bz,
						nz = ax * by - ay * bx;
					const len = Math.sqrt(nx * nx + ny * ny + nz * nz) || 1;
					nx /= len;
					ny /= len;
					nz /= len;
					// the sketch's rule: a petal's front faces the camera, so the inside of a curl
					// is lit like the outside of it
					if (nz < 0) {
						nx = -nx;
						ny = -ny;
						nz = -nz;
					}
					cb(idx, A, CAM, k, v, { x: nx, y: ny, z: nz });
				}
			}
		}
	}

	eachSample(false, (_idx, _local, cam) => {
		const sx = Math.round(SMC + (cam.x * E1x + cam.z * E1z) * SMS),
			sy = Math.round(SMC - (cam.x * E2x + cam.y * E2y + cam.z * E2z) * SMS);
		const d = cam.x * Lx + cam.y * Ly + cam.z * Lz;
		for (let dy = 0; dy < 2; dy++)
			for (let dx = 0; dx < 2; dx++) {
				const j = (sy + dy) * SMW + sx + dx;
				if (j >= 0 && j < SM.length && SM[j] < d) SM[j] = d;
			}
	});

	const TAPS = [0, 0, 5, 0, -5, 0, 0, 5, 0, -5, 11, 0, -11, 0, 0, 11, 0, -11, 8, 8, -8, 8, 8, -8, -8, -8];
	function litFrac(cam: Point) {
		const sx = Math.round(SMC + (cam.x * E1x + cam.z * E1z) * SMS),
			sy = Math.round(SMC - (cam.x * E2x + cam.y * E2y + cam.z * E2z) * SMS);
		const d = cam.x * Lx + cam.y * Ly + cam.z * Lz + 0.03;
		let n = 0;
		for (let q = 0; q < 26; q += 2) {
			const j = (sy + TAPS[q + 1]) * SMW + sx + TAPS[q];
			if (j >= 0 && j < SM.length && SM[j] <= d) n++;
		}
		return n / 13;
	}

	// ---- 2. the depth buffer, with the flower-local point of whatever won each buffer pixel
	const ZB = new Float32Array(BW * BH).fill(-1e9),
		ID = new Int32Array(BW * BH).fill(-1);

	// The back of the bloom needs its own buffer, or turning the flower round reveals nothing: the
	// sketch only ever recorded what its fixed camera could see. Both views ride the same pass, so
	// the second one costs a rotation and a projection, not another three million pos() calls.
	const ZB2 = new Float32Array(BW * BH).fill(-1e9),
		ID2 = new Int32Array(BW * BH).fill(-1);
	const BACK: Point = { x: 0, y: 0, z: 0 };

	// Only the winner's sample index is kept — one integer a pixel. Its position and normal are
	// regenerated when the rules ask for them, which happens for a few thousand cells, not for
	// every sample; writing seven float arrays per sample was most of the build's cost.
	eachSample(true, (idx, local, cam) => {
		const ix = Math.floor((BX + cam.x * ROSE_SCALE) / STEP),
			iy = Math.floor((BY - cam.y * ROSE_SCALE) / STEP);
		if (ix < 0 || iy < 0 || ix >= BW - 1 || iy >= BH - 1) return;
		for (let dy = 0; dy < 2; dy++)
			for (let dx = 0; dx < 2; dx++) {
				const j = (iy + dy) * BW + ix + dx;
				if (ZB[j] < cam.z) {
					ZB[j] = cam.z;
					ID[j] = idx;
				}
			}
		// the same sample, the flower turned half a turn: x and z negate, y is the axis
		BACK.x = -local.x;
		BACK.y = local.y;
		BACK.z = -local.z;
		toCamera(BACK, CAM2);
		const ix2 = Math.floor((BX + CAM2.x * ROSE_SCALE) / STEP),
			iy2 = Math.floor((BY - CAM2.y * ROSE_SCALE) / STEP);
		if (ix2 < 0 || iy2 < 0 || ix2 >= BW - 1 || iy2 >= BH - 1) return;
		for (let dy = 0; dy < 2; dy++)
			for (let dx = 0; dx < 2; dx++) {
				const j = (iy2 + dy) * BW + ix2 + dx;
				if (ZB2[j] < CAM2.z) {
					ZB2[j] = CAM2.z;
					ID2[j] = idx;
				}
			}
	});

	const PER = (NU + 1) * (NV + 1);
	const petalOf = (idx: number) => Math.floor(idx / PER);
	// how far along its petal a sample sits, from its index
	function vOf(idx: number) {
		const k = petalOf(idx);
		const rem = idx - k * PER;
		const iu = Math.floor(rem / (NV + 1));
		const iv = rem - iu * (NV + 1);
		const p = petals[k];
		const u = -1 + (2 * iu) / NU;
		const vmax = 1 - 0.3 * u * u + 0.04 * Math.sin(u * 3 + p.ph);
		return (iv / NV) * vmax;
	}
	const SAMPLE: Point = { x: 0, y: 0, z: 0 };
	const SAMPLE_N: Point = { x: 0, y: 0, z: 0 };
	const SAMPLE_CAM: Point = { x: 0, y: 0, z: 0 };
	// rebuild a sample from its index: the same pos() the sampler used
	function sampleAt(idx: number, out: Point, outNormal: Point) {
		const k = petalOf(idx);
		const rem = idx - k * PER;
		const iu = Math.floor(rem / (NV + 1));
		const iv = rem - iu * (NV + 1);
		const p = petals[k];
		const u = -1 + (2 * iu) / NU;
		const vmax = 1 - 0.3 * u * u + 0.04 * Math.sin(u * 3 + p.ph);
		const v = (iv / NV) * vmax;
		pos(p, u, v, out);
		pos(p, u + 0.01, v, B);
		pos(p, u, v + 0.01, C);
		const ax = B.x - out.x,
			ay = B.y - out.y,
			az = B.z - out.z,
			bx = C.x - out.x,
			by = C.y - out.y,
			bz = C.z - out.z;
		let nx = ay * bz - az * by,
			ny = az * bx - ax * bz,
			nz = ax * by - ay * bx;
		const len = Math.sqrt(nx * nx + ny * ny + nz * nz) || 1;
		nx /= len;
		ny /= len;
		nz /= len;
		if (nz < 0) {
			nx = -nx;
			ny = -ny;
			nz = -nz;
		}
		outNormal.x = nx;
		outNormal.y = ny;
		outNormal.z = nz;
	}

	// The buffer is sampled less densely than the sketch sampled it, so a few pixels inside the
	// bloom stay empty; left alone they read as silhouette to the rules below, which turns quiet
	// petals into a rash of crosses. Fill each hole from its deepest filled neighbour.
	function fillHoles(zb: Float32Array, id: Int32Array) {
		const holes: number[] = [],
			sources: number[] = [];
		for (let pass = 0; pass < 2; pass++) {
			holes.length = 0;
			sources.length = 0;
			for (let y = 1; y < BH - 1; y++)
				for (let x = 1; x < BW - 1; x++) {
					const j = y * BW + x;
					if (id[j] >= 0) continue;
					let best = -1;
					for (let d = 0; d < 4; d++) {
						const jj = y * BW + x + D4x[d] + D4y[d] * BW;
						if (id[jj] >= 0 && (best < 0 || zb[jj] > zb[best])) best = jj;
					}
					if (best < 0) continue;
					holes.push(j);
					sources.push(best);
				}
			// applied after the scan: filling in place would flood the empty background from the
			// bloom's edge, one pixel per pass, until the whole frame read as flower
			for (let i = 0; i < holes.length; i++) {
				const j = holes[i],
					best = sources[i];
				zb[j] = zb[best];
				id[j] = id[best];
			}
		}
	}
	fillHoles(ZB, ID);

	function at(artX: number, artY: number) {
		const bx = Math.floor(artX / STEP),
			by = Math.floor(artY / STEP);
		return bx < 0 || by < 0 || bx >= BW || by >= BH ? -1 : by * BW + bx;
	}

	// the lowest covered pixel near the middle is where the stem begins
	let stemY = 0,
		stemX0 = BX;
	for (let yy = BH - 1; yy > 0 && !stemY; yy--)
		for (let xx = Math.floor((BX - 80) / STEP); xx < Math.floor((BX + 80) / STEP); xx++)
			if (ID[yy * BW + xx] >= 0) {
				stemY = yy * STEP;
				stemX0 = xx * STEP;
				break;
			}

	// ---- 3. marks: one per six-pixel cell, first match wins, each carrying its own 3D point
	const MK: RoseMark[] = [];

	function runRules(zb: Float32Array, id: Int32Array, withStem: boolean) {
		function rec(t: MarkKind, artX: number, artY: number, n: number, al: number) {
			const b = t === 3 ? 16 : Math.round(clamp(al) * 15);
			if (b < 1) return;
			const j = at(artX, artY);
			if (j < 0 || id[j] < 0) return;
			const idx = id[j];
			sampleAt(idx, SAMPLE, SAMPLE_N);
			const k = petalOf(idx);
			const v = vOf(idx);
			// the petal swings about a horizontal axis across its own radius: that is the motion
			// the sketch's hinge produced, and it opens the bloom outward the same way in 3D
			const phi = petals[k].phi;
			toCamera(SAMPLE, SAMPLE_CAM);
			MK.push({
				t,
				lit: litFrac(SAMPLE_CAM),
				noise: 0.9 + 0.2 * vnoise(artX * 0.05, artY * 0.05),
				x: SAMPLE.x,
				y: SAMPLE.y,
				z: SAMPLE.z,
				nx: SAMPLE_N.x,
				ny: SAMPLE_N.y,
				nz: SAMPLE_N.z,
				n: n / ROSE_SCALE,
				b,
				petal: k,
				ax: Math.cos(phi),
				ay: 0,
				az: -Math.sin(phi),
				spread: (0.2 + (0.8 * k) / (N - 1)) * Math.pow(Math.min(1, v), 1.2)
			});
		}
		const plus = (x: number, y: number, n: number, al: number) => rec(0, x, y, n, al);
		const dash = (x: number, y: number, len: number, al: number) => rec(1, x, y, len / 2, al);
		const dot = (x: number, y: number, al: number) => rec(2, x, y, 0.7, al);
		const bold = (x: number, y: number, n: number) => rec(3, x, y, n, 1);

		const D4 = [
			[1, 0],
			[-1, 0],
			[0, 1],
			[0, -1]
		];
		const hi: number[][] = [];
		const stemEnd = stemY + STEM_LENGTH * (ROSE_H - stemY);
		for (let row = 0; row * CELL < ROSE_H; row++)
			for (let col = 0; col * CELL < ROSE_W; col++) {
				const px = col * CELL + CELL / 2 + (row % 2) * (CELL / 2),
					py = row * CELL + CELL / 2,
					h = hash(col, row),
					i0 = at(px, py);
				if (i0 < 0) continue;
				if (id[i0] < 0) {
					// outside the bloom: a stem of crosses on the axis, dashes beside it
					if (withStem && py > stemY - 10 && py < stemEnd) {
						const sxn = stemX0 - 0.12 * (py - stemY) + 9 * Math.sin((py - stemY) / 60);
						const lx = (px - BX) / ROSE_SCALE,
							ly = (BY - py) / ROSE_SCALE;
						if (Math.abs(px - sxn) < CELL / 2 + 0.01) {
							const b = Math.round(clamp(0.7 + 0.3 * Math.sin(py * 0.07)) * 15);
							if (b >= 1)
								MK.push({ t: 0, lit: 1, noise: 1, x: lx, y: ly, z: 0, nx: 0, ny: 0, nz: 1, n: 2.6 / ROSE_SCALE, b, petal: 0, ax: 0, ay: 1, az: 0, spread: 0 });
						} else if (px < sxn && px > sxn - CELL * 1.6 && row % 3 === 0) {
							const b = Math.round(0.35 * 15);
							if (b >= 1)
								MK.push({ t: 1, lit: 1, noise: 1, x: lx, y: ly, z: 0, nx: 0, ny: 0, nz: 1, n: (CELL * 1.2) / 2 / ROSE_SCALE, b, petal: 0, ax: 0, ay: 1, az: 0, spread: 0 });
						}
					}
					continue;
				}
				let sil = false,
					lip = false,
					occl = false;
				const zi = zb[i0];
				for (let d = 0; d < 4; d++) {
					const j = at(px + D4[d][0] * CELL * 0.9, py + D4[d][1] * CELL * 0.9);
					if (j < 0 || id[j] < 0) sil = true;
					else if (petalOf(id[j]) !== petalOf(id[i0])) {
						const dz = zi - zb[j];
						if (dz > 0.035) lip = true;
						else if (dz < -0.035) occl = true;
					}
				}
				sampleAt(id[i0], SAMPLE, SAMPLE_N);
				toCamera(SAMPLE, SAMPLE_CAM);
				const nl = SAMPLE_N.x * Lx + SAMPLE_N.y * Ly + SAMPLE_N.z * Lz,
					lam = Math.pow(clamp((nl + 0.45) / 1.45), 1.2) * (0.9 + 0.2 * vnoise(px * 0.05, py * 0.05)),
					fill = Math.max(0, SAMPLE_N.x * 0.5 + SAMPLE_N.y * 0.2 + SAMPLE_N.z * 0.84);
				const lit = litFrac(SAMPLE_CAM),
					rim = Math.pow(1 - SAMPLE_N.z, 3) * 0.2;
				let L = clamp(0.08 + 0.46 * lam * (0.4 + 0.6 * lit) + 0.1 * fill + rim);
				if (occl) L *= 0.6;
				if (sil) {
					plus(px, py, 2.6, 0.5 + 0.5 * Math.max(L, 0.35));
					continue;
				}
				if (lip) {
					plus(px, py, 2.8, 0.55 + 0.45 * L);
					if (L > 0.55) hi.push([px, py]);
					continue;
				}
				if (L > 0.6) {
					hi.push([px, py]);
					dash(px, py, CELL * 3.4, 1);
					continue;
				}
				if (L > 0.55) {
					dash(px, py, CELL * (1.2 + 2.2 * L), 0.55 + 0.45 * L);
					continue;
				}
				if (L > 0.42) {
					dash(px, py, CELL * (0.9 + 1.4 * L), 0.3 + 0.7 * L);
					continue;
				}
				if (L > 0.26) {
					dash(px, py, CELL * 0.8, 0.25 + 0.8 * L);
					if (h > 0.6) dot(px + CELL * 0.4, py + CELL * 0.45, 0.4);
					continue;
				}
				if (L > 0.1 && h < L * 3.2) dot(px, py, 0.2 + L);
			}
		for (let q = 0; q < hi.length; q++) if (hash(q, 3) > 0.4) bold(hi[q][0], hi[q][1], 4.4);
	}

	runRules(ZB, ID, true);
	fillHoles(ZB2, ID2);
	runRules(ZB2, ID2, false);

	return { count: MK.length, marks: MK };
}
