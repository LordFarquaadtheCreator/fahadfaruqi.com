export function reveal(node: HTMLElement) {
	node.dataset.revealed = 'true';

	if (typeof window === 'undefined') {
		return;
	}

	const reduceMotion =
		typeof window.matchMedia === 'function' &&
		window.matchMedia('(prefers-reduced-motion: reduce)').matches;
	if (reduceMotion || !('IntersectionObserver' in window)) {
		return;
	}

	node.dataset.revealed = 'false';

	const observer = new IntersectionObserver(
		(entries) => {
			if (entries.some((entry) => entry.isIntersecting)) {
				node.dataset.revealed = 'true';
				observer.disconnect();
			}
		},
		{
			threshold: 0.18,
			rootMargin: '0px 0px -48px 0px'
		}
	);

	observer.observe(node);

	return {
		destroy() {
			observer.disconnect();
		}
	};
}
