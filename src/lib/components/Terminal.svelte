<script lang="ts">
  import { terminal } from '$lib/stores/terminalStore';
  import History   from './History.svelte';
  import InputRow  from './InputRow.svelte';
  import HintBar   from './HintBar.svelte';

  let inputRow: InputRow;

  function handleBodyClick(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      inputRow?.focus();
    }
  }

  function handleTerminalKey(e: KeyboardEvent) {
    if (e.key === 'Enter' || e.key === ' ') {
      inputRow?.focus();
    }
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_noninteractive_element_interactions a11y_no_noninteractive_tabindex -->
<div
  class="terminal {$terminal.theme}"
  onkeydown={handleTerminalKey}
  tabindex="0"
  role="application"
  aria-label="terminal"
>
  <h2 class="sr-only">Portfolio terminal — navigate using commands or hint buttons</h2>

  <div class="titlebar">
    <div class="dots">
      <span class="dot" style="background:#cc241d"></span>
      <span class="dot" style="background:#d79921"></span>
      <span class="dot" style="background:#98971a"></span>
    </div>
    <span class="title-text">user@portfolio — {$terminal.path}</span>
  </div>

  <!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
  <div class="body" role="region" aria-label="terminal body" onclick={handleBodyClick}>
    <History entries={$terminal.history} />
    <InputRow bind:this={inputRow} path={$terminal.path} />
    <HintBar hints={$terminal.hints} candidates={$terminal.completionCandidates} />
  </div>
</div>

<style>
  .terminal {
    width: 100%;
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    border: 1px solid var(--gb-border);
    background: var(--gb-bg);
    font-family: 'JetBrains Mono', monospace;
    font-size: 13px;
    cursor: text;
  }
  .titlebar {
    background: var(--gb-bg-hard);
    border-bottom: 1px solid var(--gb-border);
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 10px 14px;
    flex-shrink: 0;
  }
  .dots { display: flex; gap: 6px; }
  .dot  { width: 12px; height: 12px; border-radius: 50%; display: inline-block; }
  .title-text { flex: 1; text-align: center; font-size: 11px; color: var(--gb-fg2); letter-spacing: .05em; }
  .body {
    padding: 14px 16px;
    display: flex;
    flex-direction: column;
    flex: 1;
    overflow-y: auto;
    min-height: 0;
  }

  @media (max-width: 640px) {
    .terminal { font-size: 14px; border: none; }
    .titlebar { padding: 8px 10px; }
    .dot { width: 10px; height: 10px; }
    .title-text { font-size: 10px; }
    .body { padding: 10px 12px; }
  }
</style>
