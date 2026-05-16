<script lang="ts">
  import type { HistoryEntry as HistoryEntryType } from '$lib/model/types';
  import OutputRenderer from './OutputRenderer.svelte';

  let { entry }: { entry: HistoryEntryType } = $props();
</script>

<div class="entry">
  {#if entry.command}
    <div class="cmd-row">
      <span class="p-user">user</span><span class="p-sep">@</span><span
        class="p-host">portfolio</span><span class="p-sep">:</span><span
        class="p-path">{entry.path}</span><span class="p-dollar"> $</span>
      <span class="cmd-text">{entry.command}</span>
    </div>
  {/if}
  <div class="output-block">
    <OutputRenderer output={entry.output} />
  </div>
</div>

<style>
  .entry { display: flex; flex-direction: column; gap: 5px; }
  .cmd-row { display: flex; flex-wrap: wrap; align-items: center; gap: 1px; line-height: 1.6; }
  .p-user   { color: var(--gb-green);  font-weight: 500; }
  .p-sep    { color: var(--gb-fg2); }
  .p-host   { color: var(--gb-aqua);   font-weight: 500; }
  .p-path   { color: var(--gb-purple); }
  .p-dollar { color: var(--gb-yellow); font-weight: 500; margin-right: 6px; }
  .cmd-text { color: var(--gb-fg); }
  .output-block {
    border-left: 2px solid var(--gb-border2);
    padding: 4px 0 4px 12px;
    margin-left: 2px;
  }
</style>
