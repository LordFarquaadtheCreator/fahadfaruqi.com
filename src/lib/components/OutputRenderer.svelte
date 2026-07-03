<script lang="ts">
  import type { OutputEntry, ListItem, CatPayload } from '$lib/model/types';

  let { output }: { output: OutputEntry } = $props();

  function isListItems(p: unknown): p is ListItem[] {
    return Array.isArray(p);
  }

  function isCatPayload(p: unknown): p is CatPayload {
    return typeof p === 'object' && p !== null && 'content' in p;
  }

  function formatPayload(p: unknown): string {
    if (typeof p === 'string') return p;
    if (p === null || p === undefined) return '';
    try {
      return JSON.stringify(p, null, 2);
    } catch {
      return String(p);
    }
  }
</script>

{#if output.type === 'list' && isListItems(output.payload)}
  <pre class="out-text list-output">
    {#each output.payload as item, i}
      <span class="list-item" class:dir={item.primary.endsWith('/')}>
        {item.primary}{#if item.meta}{' '}{item.meta}{/if}
      </span>{#if i < output.payload.length - 1}{'  '}{/if}
    {/each}
  </pre>

{:else if output.type === 'text'}
  <pre class="out-text">{formatPayload(output.payload)}</pre>

{:else if output.type === 'error'}
  <span class="out-error">{formatPayload(output.payload)}</span>

{:else if output.type === 'success'}
  <span class="out-success">{formatPayload(output.payload)}</span>

{:else if output.type === 'warning'}
  <span class="out-warning">{formatPayload(output.payload)}</span>

{:else if output.type === 'cat' && isCatPayload(output.payload)}
  <pre class="out-text">{output.payload.content}</pre>

{:else if output.type === 'json' && typeof output.payload === 'object'}
  <pre class="out-text">{JSON.stringify(output.payload, null, 2)}</pre>

{:else}
  <pre class="out-text">{formatPayload(output.payload)}</pre>
{/if}

<style>
  .list-output { white-space: pre-wrap; word-break: break-word; }
  .list-item { display: inline; }
  .list-item.dir { color: var(--gb-aqua); }
  .out-error   { color: var(--gb-red); }
  .out-success { color: var(--gb-green); }
  .out-warning { color: var(--gb-yellow); }
  .out-text    { color: var(--gb-fg2); white-space: pre-wrap; font-size: 12px; line-height: 1.8; font-family: inherit; }
</style>
