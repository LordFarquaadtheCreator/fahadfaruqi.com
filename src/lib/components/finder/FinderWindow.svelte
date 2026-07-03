<script lang="ts">
	import type { WindowVM } from '$lib/finder/finderStore.svelte';
	import FinderToolbar from './FinderToolbar.svelte';
	import FinderSidebar from './FinderSidebar.svelte';
	import IconGrid from './content/IconGrid.svelte';
	import ListTable from './content/ListTable.svelte';
	import ColumnView from './content/ColumnView.svelte';
	import GalleryView from './content/GalleryView.svelte';
	import MarkdownPane from './content/MarkdownPane.svelte';

	interface Props {
		win: WindowVM;
	}

	let { win }: Props = $props();

	let dragging = $state(false);
	let dragStartX = 0;
	let dragStartY = 0;

	function onPointerDown(e: PointerEvent) {
		const target = e.target as HTMLElement;
		if (!target.closest('.title-bar')) return;
		dragging = true;
		dragStartX = e.clientX;
		dragStartY = e.clientY;
		(target as HTMLElement).setPointerCapture(e.pointerId);
		window.os.focus(win.id);
	}

	function onPointerMove(e: PointerEvent) {
		if (!dragging) return;
		const dx = e.clientX - dragStartX;
		const dy = e.clientY - dragStartY;
		window.os.drag(win.id, dx, dy);
		dragStartX = e.clientX;
		dragStartY = e.clientY;
	}

	function onPointerUp() {
		dragging = false;
	}

	function onWindowClick() {
		window.os.focus(win.id);
	}

	function onCloseClick(e: MouseEvent) {
		e.stopPropagation();
		window.os.close(win.id);
	}
</script>

<div
	class="finder-window"
	role="dialog"
	aria-label="Finder window {win.title}"
	tabindex="-1"
	style:left="{win.x}px"
	style:top="{win.y}px"
	style:z-index={win.z}
	style:width="{win.width}px"
	style:height="{win.height}px"
	onclick={onWindowClick}
	onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); onWindowClick(); } }}
>
	<div
		class="title-bar"
		role="button"
		aria-label="Drag to move"
		tabindex="-1"
		onpointerdown={onPointerDown}
		onpointermove={onPointerMove}
		onpointerup={onPointerUp}
	>
		<div class="traffic-lights">
			<button class="tl close" onclick={onCloseClick} aria-label="Close"></button>
			<button class="tl minimize" aria-label="Minimize" disabled></button>
			<button class="tl maximize" aria-label="Maximize" disabled></button>
		</div>
		<span class="title">{win.title}</span>
	</div>

	<FinderToolbar {win} />

	<div class="body">
		<FinderSidebar {win} />
		<div class="content">
			{#if win.content.kind === 'folder'}
				{#if win.viewMode === 'icon'}
					<IconGrid entries={win.content.entries ?? []} winId={win.id} selected={win.selectedEntry} />
				{:else if win.viewMode === 'list'}
					<ListTable entries={win.content.entries ?? []} winId={win.id} selected={win.selectedEntry} />
				{:else if win.viewMode === 'column'}
					<ColumnView entries={win.content.entries ?? []} winId={win.id} selected={win.selectedEntry} />
				{:else if win.viewMode === 'gallery'}
					<GalleryView entries={win.content.entries ?? []} winId={win.id} selected={win.selectedEntry} />
				{/if}
			{:else if win.content.kind === 'markdown'}
				<MarkdownPane html={win.content.html ?? ''} />
			{/if}
		</div>
	</div>
</div>

<style>
	.finder-window {
		position: absolute;
		display: flex;
		flex-direction: column;
		background: var(--gb-bg, #282828);
		border: 1px solid var(--gb-border, #504945);
		box-shadow: 0 10px 40px rgba(0, 0, 0, 0.5);
		overflow: hidden;
		border-radius: 10px;
		font-family: 'JetBrains Mono', monospace;
		font-size: 13px;
		color: var(--gb-fg, #ebdbb2);
	}

	.title-bar {
		display: flex;
		align-items: center;
		gap: 8px;
		padding: 8px 12px;
		background: var(--gb-bg-hard, #1d2021);
		border-bottom: 1px solid var(--gb-border, #504945);
		cursor: default;
		user-select: none;
		-webkit-user-select: none;
	}

	.traffic-lights {
		display: flex;
		gap: 6px;
		flex-shrink: 0;
	}

	.tl {
		width: 12px;
		height: 12px;
		border-radius: 50%;
		border: none;
		padding: 0;
		cursor: pointer;
	}

	.tl:disabled {
		cursor: default;
		opacity: 0.5;
	}

	.tl.close {
		background: #ff5f56;
	}
	.tl.minimize {
		background: #ffbd2e;
	}
	.tl.maximize {
		background: #27c93f;
	}

	.title {
		flex: 1;
		text-align: center;
		font-size: 12px;
		color: var(--gb-fg2, #a89984);
		margin-right: 52px;
	}

	.body {
		display: flex;
		flex: 1;
		overflow: hidden;
	}

	.content {
		flex: 1;
		overflow: auto;
		padding: 8px;
	}
</style>
