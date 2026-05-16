<script lang="ts">
  import type { HistoryEntry as HistoryEntryType } from '$lib/model/types';
  import HistoryEntryComponent from './HistoryEntry.svelte';
  import { tick } from 'svelte';

  let { entries }: { entries: HistoryEntryType[] } = $props();

  let scrollEl: HTMLDivElement;

  $effect(() => {
    entries;
    tick().then(() => {
      if (scrollEl) scrollEl.scrollTop = scrollEl.scrollHeight;
    });
  });
</script>

<div class="history" bind:this={scrollEl}>
  {#each entries as entry (entry.id)}
    <HistoryEntryComponent {entry} />
  {/each}
</div>

<style>
  .history {
    flex: 1;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    gap: 10px;
    padding-bottom: 8px;
    scrollbar-width: thin;
    scrollbar-color: var(--gb-border) transparent;
  }
</style>
