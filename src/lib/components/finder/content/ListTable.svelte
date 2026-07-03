<script lang="ts">
	import type { EntryVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		entries: EntryVM[];
		winId: string;
		selected: string;
	}

	let { entries, winId, selected }: Props = $props();
</script>

<table class="list-table">
	<thead>
		<tr>
			<th>Name</th>
			<th>Kind</th>
			<th>Size</th>
		</tr>
	</thead>
	<tbody>
		{#each entries as entry}
			<tr
				class:selected={selected === entry.name}
				onclick={() => window.os.selectEntry(winId, entry.name)}
				ondblclick={() => window.os.openEntry(winId, entry.path)}
			>
				<td>{entry.kind === 'dir' ? '📁' : '📄'} {entry.name}</td>
				<td>{entry.kind}</td>
				<td>{entry.size}</td>
			</tr>
		{/each}
	</tbody>
</table>

<style>
	.list-table {
		width: 100%;
		border-collapse: collapse;
		font-size: 12px;
	}

	.list-table th {
		text-align: left;
		padding: 4px 8px;
		border-bottom: 1px solid var(--gb-border, #504945);
		color: var(--gb-fg2, #a89984);
		font-weight: 500;
	}

	.list-table td {
		padding: 4px 8px;
		cursor: pointer;
	}

	.list-table tr:hover {
		background: var(--gb-border2, #3c3836);
	}

	.list-table tr.selected {
		background: var(--gb-border, #504945);
	}
</style>
