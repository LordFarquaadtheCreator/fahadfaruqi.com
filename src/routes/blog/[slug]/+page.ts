import { blogPosts, getBlogPost } from '$lib/blogs';
import { error } from '@sveltejs/kit';

export const entries = () => blogPosts.map((post) => ({ slug: post.slug }));

export function load({ params }) {
	const post = getBlogPost(params.slug);

	if (!post) {
		error(404, 'Post not found');
	}

	const index = blogPosts.findIndex((entry) => entry.slug === post.slug);
	const nextPost = blogPosts[(index + 1) % blogPosts.length];

	return {
		post,
		nextPost
	};
}
