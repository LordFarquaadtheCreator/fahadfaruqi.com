// The rose's geometry: 44 petal shells wrapped around a vertical axis, sampled into a cloud of
// points that each carry their own surface normal. What mark a point becomes, and how bright it is,
// is decided per frame in the renderer — in 3D that depends on where the camera and the light are,
// not on one fixed view.

export interface Rose {
	count: number;
	/** Rose-local space: x right, y up the stem, z toward the viewer at yaw 0. */
	x: Float32Array;
	y: Float32Array;
	z: Float32Array;
	nx: Float32Array;
	ny: Float32Array;
	nz: Float32Array;
	/** The petal the point belongs to: the wind swings each petal as one surface. */
	petal: Int32Array;
}

/** How many petals the bloom has. The wind loop counts them. */
export const PETALS = 44;
// Samples across each petal and along it: the surface density the old six-pixel screen grid gave,
// but now on the surface itself, so the back of the bloom is as detailed as the front.
const ACROSS = 14;
const ALONG = 26;

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
	// The draw order is the sketch's own, so a seed still names the flower it always did.
	const R = rng(seed);
	const petals: Petal[] = [];
	for (let i = 0; i < PETALS; i++) {
		const t = i / (PETALS - 1),
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

	// One petal's surface at (u across, v along). No camera in here: the flower is built in its own
	// frame and the renderer supplies the view.
	function pos(p: Petal, u: number, v: number, o: Point) {
		const vv = Math.max(v, 0),
			w = p.w * (0.5 + 0.5 * Math.sqrt(vv)),
			phi = p.phi + u * w + 0.25 * p.tw * v;
		let k = p.b1 - p.b0;
		if (Math.abs(k) < 1e-3) k = 1e-3;
		const bb = p.b0 + k * vv; // angle from vertical, grows along the petal
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

	const total = PETALS * ACROSS * ALONG;
	const rose: Rose = {
		count: total,
		x: new Float32Array(total),
		y: new Float32Array(total),
		z: new Float32Array(total),
		nx: new Float32Array(total),
		ny: new Float32Array(total),
		nz: new Float32Array(total),
		petal: new Int32Array(total)
	};

	const A: Point = { x: 0, y: 0, z: 0 };
	const B: Point = { x: 0, y: 0, z: 0 };
	const C: Point = { x: 0, y: 0, z: 0 };
	let n = 0;
	for (let k = 0; k < PETALS; k++) {
		const p = petals[k];
		for (let iu = 0; iu < ACROSS; iu++) {
			const u = -1 + (2 * iu) / (ACROSS - 1);
			const vmax = 1 - 0.3 * u * u + 0.04 * Math.sin(u * 3 + p.ph);
			for (let iv = 0; iv < ALONG; iv++) {
				const v = (iv / (ALONG - 1)) * vmax;
				pos(p, u, v, A);
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
				// point it away from the stem, so every petal's front face is its outside
				if (nx * A.x + nz * A.z < 0) {
					nx = -nx;
					ny = -ny;
					nz = -nz;
				}
				rose.x[n] = A.x;
				rose.y[n] = A.y;
				rose.z[n] = A.z;
				rose.nx[n] = nx;
				rose.ny[n] = ny;
				rose.nz[n] = nz;
				rose.petal[n] = k;
				n++;
			}
		}
	}
	return rose;
}
