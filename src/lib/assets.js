// Images, icons and wallpapers are not part of this repo's build: they are served
// from Cloudflare (R2 bucket `assets`, custom domain assets.fahadfaruqi.com) under
// the www/ prefix. The production origin below is the default; VITE_ASSETS_BASE
// overrides it at build time, and window.__ASSETS_BASE__ (set in app.html) carries
// the same default for code that runs before this module.
export const ASSETS_BASE =
	import.meta.env.VITE_ASSETS_BASE ??
	/** @type {any} */ (globalThis).__ASSETS_BASE__ ??
	'https://assets.fahadfaruqi.com/www';

/**
 * assetUrl turns an absolute asset path such as "/icons/finder.png" into its URL.
 * @param {string} path
 * @returns {string}
 */
export function assetUrl(path) {
	return `${ASSETS_BASE}${path}`;
}
