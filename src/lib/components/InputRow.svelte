<script lang="ts">
  import { terminal } from '$lib/stores/terminalStore';

  let { path }: { path: string } = $props();

  let value = $state('');
  let inputEl: HTMLInputElement;

  export function focus() { inputEl?.focus(); }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Enter') {
      terminal.execute(value);
      value = '';
      return;
    }

    if (e.key === 'Tab') {
      e.preventDefault();
      value = terminal.tabComplete(value);
      return;
    }

    if (e.key === 'ArrowUp') {
      e.preventDefault();
      value = terminal.historyUp();
      return;
    }

    if (e.key === 'ArrowDown') {
      e.preventDefault();
      value = terminal.historyDown();
    }
  }
</script>

<div class="input-row">
  <span class="p-user">user</span><span class="p-sep">@</span><span
    class="p-host">portfolio</span><span class="p-sep">:</span><span
    class="p-path">{path}</span><span class="p-dollar"> $</span>
  <input
    bind:this={inputEl}
    bind:value
    onkeydown={handleKeydown}
    class="cmd-input"
    autocomplete="off"
    spellcheck="false"
    aria-label="terminal input"
  />
</div>

<style>
  .input-row {
    display: flex;
    align-items: center;
    border-top: 1px solid var(--gb-border2);
    padding-top: 10px;
    flex-wrap: wrap;
    gap: 1px;
  }
  .p-user   { color: var(--gb-green);  font-weight: 500; }
  .p-sep    { color: var(--gb-fg2); }
  .p-host   { color: var(--gb-aqua);   font-weight: 500; }
  .p-path   { color: var(--gb-purple); }
  .p-dollar { color: var(--gb-yellow); font-weight: 500; margin-right: 4px; }
  .cmd-input {
    background: transparent;
    border: none;
    outline: none;
    color: var(--gb-fg);
    font-family: inherit;
    font-size: inherit;
    flex: 1;
    caret-color: var(--gb-yellow);
    min-width: 40px;
  }

  @media (max-width: 640px) {
    .input-row { padding-top: 8px; }
    .cmd-input { font-size: 16px; }
  }
</style>
