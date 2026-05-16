<script lang="ts">
  import type { OutputEntry, ListItem } from '$lib/model/types';

  let { output }: { output: OutputEntry } = $props();

  function isListItems(p: unknown): p is ListItem[] {
    return Array.isArray(p);
  }
</script>

{#if output.type === 'list' && isListItems(output.payload)}
  <div class="list-output">
    {#each output.payload as item}
      <div class="list-row">
        <span class="list-primary">{item.primary}</span>
        {#if item.meta}<span class="list-meta">{item.meta}</span>{/if}
        {#if item.secondary}<div class="list-secondary">{item.secondary}</div>{/if}
      </div>
    {/each}
  </div>

{:else if output.type === 'error'}
  <span class="out-error">✗  {output.payload}</span>

{:else if output.type === 'success'}
  <span class="out-success">→ {output.payload}</span>

{:else if output.type === 'warning'}
  <span class="out-warning">{output.payload}</span>

{:else}
  <pre class="out-text">{output.payload}</pre>
{/if}

<style>
  .list-output { display: flex; flex-direction: column; gap: 8px; }
  .list-row { display: flex; flex-direction: column; gap: 2px; }
  .list-primary { color: var(--gb-aqua); font-weight: 500; }
  .list-secondary { color: var(--gb-fg); font-size: 12px; line-height: 1.6; }
  .list-meta { color: var(--gb-purple); font-size: 11px; margin-left: auto; }
  .out-error   { color: var(--gb-red); }
  .out-success { color: var(--gb-green); }
  .out-warning { color: var(--gb-yellow); }
  .out-text    { color: var(--gb-fg2); white-space: pre-wrap; font-size: 12px; line-height: 1.8; font-family: inherit; }
</style>
