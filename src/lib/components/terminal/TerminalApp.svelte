<script lang="ts">
  import type { WindowVM, TerminalContentVM, ShellResponse } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { window: win, content }: { window: WindowVM; content: TerminalContentVM } = $props();

  let input = $state('');
  let scrollRef = $state<HTMLDivElement | null>(null);

  function runCommand() {
    if (!input.trim()) return;
    const output = os()?.terminalRun?.(win.id, input) ?? '';
    input = '';
    requestAnimationFrame(() => {
      if (scrollRef) scrollRef.scrollTop = scrollRef.scrollHeight;
    });
  }

  function onKeydown(e: KeyboardEvent) {
    if (e.key === 'Enter') {
      e.preventDefault();
      runCommand();
    } else if (e.key === 'Tab') {
      e.preventDefault();
      const result = os()?.terminalComplete?.(win.id, input);
      if (result) {
        try {
          const candidates = JSON.parse(result) as string[];
          if (candidates.length === 1) {
            const parts = input.split(' ');
            parts[parts.length - 1] = candidates[0];
            input = parts.join(' ');
          }
        } catch {}
      }
    }
  }
</script>

<div class="terminal-app" bind:this={scrollRef}>
  <div class="terminal-output">
    {#each content.history as entry}
      <div class="terminal-line">
        <span class="terminal-prompt">$</span>
        <span class="terminal-input">{entry.input}</span>
      </div>
      <div class="terminal-result">{entry.output}</div>
    {/each}
  </div>
  <div class="terminal-input-row">
    <span class="terminal-prompt">$</span>
    <input
      type="text"
      bind:value={input}
      onkeydown={onKeydown}
      autocomplete="off"
      spellcheck="false"
    />
  </div>
</div>
