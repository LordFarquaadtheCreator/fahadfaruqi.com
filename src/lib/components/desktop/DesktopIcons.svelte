<script lang="ts">
  import type { DesktopIconVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { icons }: { icons: DesktopIconVM[] } = $props();

  function onDoubleClick(path: string) {
    os()?.openFile?.(path);
  }
</script>

<div class="desktop-icons">
  {#each icons as icon}
    <button
      class="desktop-icon"
      style="right: {icon.x}px; top: {icon.y}px;"
      ondblclick={() => onDoubleClick(icon.path)}
    >
      <img class="icon-img" src={icon.icon} alt={icon.name} />
      <span class="icon-label">{icon.name}</span>
    </button>
  {/each}
</div>
