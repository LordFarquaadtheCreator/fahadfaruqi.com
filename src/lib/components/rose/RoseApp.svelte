<script lang="ts">
  import type { RoseContentVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';
  import { mountRose } from '$lib/rose/mount';

  let { content }: { content: RoseContentVM } = $props();

  let canvas: HTMLCanvasElement | undefined = $state();
  let drawn = $state(false);

  // The seed is the app's state and Go owns it, so a shuffle arrives as new props: the drawing is
  // torn down and rebuilt from the new seed rather than patched.
  $effect(() => {
    const seed = content.seed;
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
  aria-label="Rose — click for a new drawing"
  onclick={() => os()?.roseShuffle?.()}
>
  <canvas bind:this={canvas} class="rose-canvas" aria-hidden="true"></canvas>
  {#if !drawn}
    <span class="rose-status">Drawing</span>
  {/if}
</button>
