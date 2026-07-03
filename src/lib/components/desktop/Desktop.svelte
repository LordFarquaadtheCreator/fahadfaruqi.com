<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { initOS, subscribe, setViewport, os } from '$lib/os/osStore.svelte';
  import type { RootVM } from '$lib/os/types';
  import Wallpaper from './Wallpaper.svelte';
  import MenuBar from './MenuBar.svelte';
  import DesktopIcons from './DesktopIcons.svelte';
  import Dock from './Dock.svelte';
  import WindowLayer from './WindowLayer.svelte';

  let vm = $state<RootVM | null>(null);

  onMount(() => {
    initOS();
    const unsub = subscribe((v) => { vm = v; });
    const onResize = () => setViewport(window.innerWidth, window.innerHeight);
    window.addEventListener('resize', onResize);
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') os()?.closeFocused?.();
    };
    window.addEventListener('keydown', onKey);
    return () => {
      unsub();
      window.removeEventListener('resize', onResize);
      window.removeEventListener('keydown', onKey);
    };
  });
</script>

<Wallpaper />
{#if vm}
  <MenuBar menuBar={vm.menuBar} />
  <DesktopIcons icons={vm.desktop.icons} />
  <WindowLayer windows={vm.windows} focusedId={vm.focusedId} />
  <Dock dock={vm.dock} />
{/if}
