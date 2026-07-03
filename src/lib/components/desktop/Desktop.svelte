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

  function isDark(): boolean {
    const theme = document.documentElement.getAttribute('data-theme') || 'system';
    if (theme === 'dark') return true;
    if (theme === 'light') return false;
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
  }

  onMount(() => {
    initOS();

    const syncTheme = () => os()?.setTheme?.(isDark());
    syncTheme();

    const mq = window.matchMedia('(prefers-color-scheme: dark)');
    const onMqChange = () => syncTheme();
    mq.addEventListener('change', onMqChange);

    const observer = new MutationObserver(() => syncTheme());
    observer.observe(document.documentElement, { attributes: true, attributeFilter: ['data-theme'] });

    const unsub = subscribe((v) => { vm = v; });
    const onResize = () => setViewport(window.innerWidth, window.innerHeight);
    window.addEventListener('resize', onResize);
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') os()?.closeFocused?.();
    };
    window.addEventListener('keydown', onKey);
    return () => {
      unsub();
      mq.removeEventListener('change', onMqChange);
      observer.disconnect();
      window.removeEventListener('resize', onResize);
      window.removeEventListener('keydown', onKey);
    };
  });
</script>

{#if vm}
  <Wallpaper src={vm.desktop.wallpaper} />
  <MenuBar menuBar={vm.menuBar} />
  <DesktopIcons icons={vm.desktop.icons} />
  <WindowLayer windows={vm.windows} focusedId={vm.focusedId} />
  <Dock dock={vm.dock} />
{/if}
