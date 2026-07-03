<script lang="ts">
  import type {
    WindowVM,
    TerminalContentVM,
    ShellResponse,
    LsPayload,
    FSEntry,
    CatPayload,
    TextPayload,
    EnvPayload,
    HistoryPayload,
    ErrorPayload,
  } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { window: win, content }: { window: WindowVM; content: TerminalContentVM } = $props();

  let input = $state('');
  let scrollRef = $state<HTMLDivElement | null>(null);
  let inputRef = $state<HTMLInputElement | null>(null);
  let historyIdx = $state(-1);
  let savedInput = '';

  function focusInput() {
    inputRef?.focus();
  }

  function runCommand() {
    if (!input.trim()) return;
    os()?.terminalRun?.(win.id, input);
    input = '';
    historyIdx = -1;
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
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      const hist = content.history;
      if (hist.length === 0) return;
      if (historyIdx === -1) {
        savedInput = input;
        historyIdx = hist.length - 1;
      } else if (historyIdx > 0) {
        historyIdx--;
      }
      input = hist[historyIdx].input;
    } else if (e.key === 'ArrowDown') {
      e.preventDefault();
      if (historyIdx === -1) return;
      const hist = content.history;
      if (historyIdx < hist.length - 1) {
        historyIdx++;
        input = hist[historyIdx].input;
      } else {
        historyIdx = -1;
        input = savedInput;
      }
    }
  }

  function parseResponse(output: string): ShellResponse | null {
    try {
      return JSON.parse(output) as ShellResponse;
    } catch {
      return null;
    }
  }

  function formatSize(bytes: number): string {
    if (bytes < 1024) return `${bytes}B`;
    if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)}K`;
    return `${(bytes / (1024 * 1024)).toFixed(1)}M`;
  }

  function shortName(path: string): string {
    const parts = path.split('/');
    return parts[parts.length - 1] || path;
  }
</script>

<div class="terminal-app" bind:this={scrollRef} onclick={focusInput} onkeydown={focusInput} role="application" tabindex="-1">
  <div class="terminal-output">
    {#if content.history.length === 0}
      <div class="terminal-welcome">
        Welcome to iTerm — type <span class="terminal-cmd">help</span> for commands.
        <br />
        Read-only filesystem. Try <span class="terminal-cmd">ls</span>, <span class="terminal-cmd">cat readme.md</span>, or <span class="terminal-cmd">open resume/</span>.
      </div>
    {/if}

    {#each content.history as entry}
      <div class="terminal-line">
        <span class="terminal-user">fahad@mac</span>
        <span class="terminal-colon">:</span>
        <span class="terminal-path">{content.cwd}</span>
        <span class="terminal-prompt">$</span>
        <span class="terminal-input">{entry.input}</span>
      </div>
      {@const resp = parseResponse(entry.output)}
      {#if resp}
        {@const payload = resp.payload as Record<string, unknown>}
        {#if resp.type === 'ls'}
          {@const ls = payload as unknown as LsPayload}
          {#if ls.entries?.length}
            <div class="terminal-ls">
              {#each ls.entries as entry_item}
                <span class="terminal-ls-{entry_item.type === 'dir' ? 'dir' : 'file'}">
                  {entry_item.name}{entry_item.type === 'dir' ? '/' : ''}
                </span>
              {/each}
            </div>
          {/if}
        {:else if resp.type === 'cat'}
          {@const cat = payload as unknown as CatPayload}
          <pre class="terminal-cat">{cat.content}</pre>
        {:else if resp.type === 'pwd'}
          {@const txt = payload as unknown as TextPayload}
          <div class="terminal-text">{txt?.text ?? resp.cwd}</div>
        {:else if resp.type === 'echo' || resp.type === 'apps'}
          {@const txt = payload as unknown as TextPayload}
          <pre class="terminal-text">{txt?.text ?? ''}</pre>
        {:else if resp.type === 'env'}
          {@const env = payload as unknown as EnvPayload}
          {#if env?.vars}
            {#each Object.entries(env.vars) as [key, val]}
              <div class="terminal-env"><span class="terminal-env-key">{key}</span>=<span class="terminal-env-val">{val}</span></div>
            {/each}
          {/if}
        {:else if resp.type === 'history'}
          {@const hist = payload as unknown as HistoryPayload}
          {#if hist?.entries}
            {#each hist.entries as cmd, i}
              <div class="terminal-text">{i + 1}  {cmd}</div>
            {/each}
          {/if}
        {:else if resp.type === 'open'}
          {@const txt = payload as unknown as TextPayload}
          <div class="terminal-text">{txt?.text ?? ''}</div>
        {:else if resp.type === 'error'}
          {@const err = payload as unknown as ErrorPayload}
          <div class="terminal-error">zsh: {err?.message ?? 'unknown error'}</div>
        {:else if resp.type === 'clear'}
          <!-- clear handled by Go, no output -->
        {:else if resp.type === 'noop' || resp.type === 'varset'}
          <!-- silent commands produce no output -->
        {:else}
          <div class="terminal-text">{entry.output}</div>
        {/if}
      {:else}
        <div class="terminal-text">{entry.output}</div>
      {/if}
    {/each}
  </div>

  <div class="terminal-input-row">
    <span class="terminal-user">fahad@mac</span>
    <span class="terminal-colon">:</span>
    <span class="terminal-path">{content.cwd}</span>
    <span class="terminal-prompt">$</span>
    <input
      type="text"
      bind:value={input}
      bind:this={inputRef}
      onkeydown={onKeydown}
      autocomplete="off"
      spellcheck="false"
      class="terminal-input-field"
    />
  </div>
</div>
