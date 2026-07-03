<script lang="ts">
  import type { WindowVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';
  import Window from '$lib/components/window/Window.svelte';
  import FinderApp from '$lib/components/finder/FinderApp.svelte';
  import TerminalApp from '$lib/components/terminal/TerminalApp.svelte';
  import ImageViewer from '$lib/components/viewer/ImageViewer.svelte';
  import TextViewer from '$lib/components/viewer/TextViewer.svelte';
  import EphemeralPreview from '$lib/components/viewer/EphemeralPreview.svelte';
  import SplashApp from '$lib/components/placeholder/SplashApp.svelte';

  let { windows, focusedId }: { windows: WindowVM[]; focusedId: string } = $props();
</script>

<div class="window-layer">
  {#each windows as win (win.id)}
    <Window
      window={win}
      focused={win.id === focusedId}
      onclose={() => os()?.close?.(win.id)}
      onfocus={() => os()?.focus?.(win.id)}
      ondrag={(dx, dy) => os()?.drag?.(win.id, dx, dy)}
      onresize={(w, h) => os()?.resize?.(win.id, w, h)}
      onminimize={() => os()?.minimize?.(win.id)}
      onmaximize={() => os()?.maximize?.(win.id)}
    >
      {#if win.content.kind === 'finder'}
        <FinderApp window={win} content={win.content} />
      {:else if win.content.kind === 'terminal'}
        <TerminalApp window={win} content={win.content} />
      {:else if win.content.kind === 'viewer-image'}
        <ImageViewer content={win.content} />
      {:else if win.content.kind === 'viewer-text'}
        <TextViewer content={win.content} />
      {:else if win.content.kind === 'ephemeral-image'}
        <EphemeralPreview content={win.content} />
      {:else if win.content.kind === 'splash'}
        <SplashApp content={win.content} />
      {/if}
    </Window>
  {/each}
</div>
