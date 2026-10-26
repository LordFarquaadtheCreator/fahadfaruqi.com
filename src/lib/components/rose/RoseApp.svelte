<script lang="ts">
  import type { RoseContentVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';
  import { mountRose } from '$lib/rose/mount';

  let { content }: { content: RoseContentVM } = $props();

  let canvas: HTMLCanvasElement | undefined = $state();
  let drawn = $state(false);

  // The seed is the app's state and Go owns it. `$derived` keeps this a number, so a VM push that
  // carries the same seed — a window drag, a focus change — does not tear the drawing down.
  const seed = $derived(content.seed);

  $effect(() => {
    const node = canvas;
    if (!node) return;
    drawn = false;
    return mountRose(node, seed, () => {
      drawn = true;
    });
  });
</script>

<button
  class="rose-app"
  type="button"
  aria-label="Rose — drag to turn it, double-click for a new drawing"
  ondblclick={() => os()?.roseShuffle?.()}
>
  <canvas bind:this={canvas} class="rose-canvas" aria-hidden="true"></canvas>
  {#if !drawn}
    <span class="rose-status">Drawing</span>
  {/if}
</button>
