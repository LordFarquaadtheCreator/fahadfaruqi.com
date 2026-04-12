import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		paths: {
			base: process.env.BASE_PATH || ''
		},
		prerender: {
			handleMissingId: 'ignore',
			handleHttpError: ({ path, message }) => {
				if (path === '/[...path]') return;
				console.warn(message);
			}
		}
	},
	vitePlugin: {
		dynamicCompileOptions: ({ filename }) => filename.includes('node_modules') ? undefined : { runes: true }
	}
};

export default config;
