<script lang="ts">
  import type { Hint } from '$lib/model/types';
  import { terminal } from '$lib/stores/terminalStore';

  let { hints, candidates = [] }: { hints: Hint[]; candidates?: string[] } = $props();
</script>

{#if candidates.length > 1}
  <div class="candidates">
    {#each candidates as c}
      <span class="candidate">{c}</span>
    {/each}
  </div>
{/if}

{#if hints.length}
  <div class="hint-bar">
    <span class="hint-label">suggestions:</span>
    <div class="pills">
      {#each hints as hint}
        <button class="pill" onclick={() => terminal.execute(hint.cmd)}>
          {hint.label}
        </button>
      {/each}
    </div>
  </div>
{/if}

<style>
  .candidates {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    padding: 6px 0;
    border-top: 1px dashed var(--gb-border2);
    margin-bottom: 4px;
  }
  .candidate { color: var(--gb-aqua); font-size: 12px; }
  .hint-bar  { padding: 8px 0 2px; }
  .hint-label { font-size: 11px; color: var(--gb-fg3); display: block; margin-bottom: 5px; }
  .pills     { display: flex; flex-wrap: wrap; gap: 6px; }
  .pill {
    background: transparent;
    border: 1px solid var(--gb-border);
    border-radius: 5px;
    color: var(--gb-fg2);
    font-family: inherit;
    font-size: 11px;
    padding: 3px 9px;
    cursor: pointer;
    transition: border-color 0.12s, color 0.12s;
  }
  .pill:hover { border-color: var(--gb-yellow); color: var(--gb-yellow); }

  @media (max-width: 640px) {
    .candidates { gap: 8px; }
    .candidate { font-size: 13px; }
    .hint-label { font-size: 10px; }
    .pill { font-size: 12px; padding: 5px 11px; }
  }
</style>
