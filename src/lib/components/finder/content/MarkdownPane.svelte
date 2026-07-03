<script lang="ts">
	interface Props {
		html: string;
	}

	let { html }: Props = $props();
</script>

<div class="markdown-pane">
	{@html html}
</div>

<style>
	.markdown-pane {
		padding: 12px;
		line-height: 1.6;
		overflow: auto;
		font-size: 13px;
	}

	.markdown-pane :global(h1) {
		color: var(--gb-fg, #ebdbb2);
		margin: 12px 0 8px;
	}
	.markdown-pane :global(h2) {
		color: var(--gb-fg, #ebdbb2);
		margin: 12px 0 8px;
	}
	.markdown-pane :global(h3) {
		color: var(--gb-fg, #ebdbb2);
		margin: 12px 0 8px;
	}

	.markdown-pane :global(p) {
		margin: 8px 0;
		color: var(--gb-fg2, #a89984);
	}

	.markdown-pane :global(a) {
		color: var(--gb-aqua, #83a598);
	}

	.markdown-pane :global(ul) {
		margin: 8px 0;
		padding-left: 20px;
	}
	.markdown-pane :global(ol) {
		margin: 8px 0;
		padding-left: 20px;
	}

	.markdown-pane :global(li) {
		margin: 2px 0;
		color: var(--gb-fg2, #a89984);
	}

	.markdown-pane :global(blockquote) {
		border-left: 3px solid var(--gb-border, #504945);
		padding-left: 12px;
		margin: 8px 0;
		color: var(--gb-fg3, #665c54);
	}

	.markdown-pane :global(img) {
		max-width: 100%;
		border-radius: 4px;
	}

	.markdown-pane :global(hr) {
		border: none;
		border-top: 1px solid var(--gb-border, #504945);
		margin: 12px 0;
	}
</style>
