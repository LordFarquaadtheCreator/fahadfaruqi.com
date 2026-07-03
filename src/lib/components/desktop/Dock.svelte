<script lang="ts">
  import type { DockVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { dock }: { dock: DockVM } = $props();
</script>

<div class="dock">
  {#each dock.items as item}
    <button
      class="dock-icon"
      class:running={item.running}
      class:active={item.active}
      onclick={() => os()?.dockClick?.(item.appId)}
    >
      <span class="dock-tooltip">{item.name}</span>
      <img src={item.icon} alt={item.name} />
      {#if item.running}
        <span class="dock-dot"></span>
      {/if}
    </button>
  {/each}
  <div class="dock-divider"></div>
  <button class="dock-icon trash">
    <span class="dock-tooltip">{dock.trash.name}</span>
    <img src={dock.trash.icon} alt={dock.trash.name} />
  </button>
</div>
