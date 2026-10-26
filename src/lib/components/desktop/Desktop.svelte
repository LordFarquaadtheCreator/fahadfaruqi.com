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
  // the menu bar inks itself to suit the wallpaper, so the theme has to be reactive, not a read of
  // the DOM at render time
  let isDarkMode = $state(false);

  function isDark(): boolean {
    const theme = document.documentElement.getAttribute('data-theme') || 'system';
    if (theme === 'dark') return true;
    if (theme === 'light') return false;
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
  }

  onMount(() => {
    initOS();

    // VITE_THEME env var overrides data-theme for dev
    const envTheme = import.meta.env.VITE_THEME as string | undefined;
    if (envTheme) {
      document.documentElement.setAttribute('data-theme', envTheme);
    }

    const syncTheme = () => {
      isDarkMode = isDark();
      os()?.setTheme?.(isDarkMode);
    };
    syncTheme();

    const mq = window.matchMedia('(prefers-color-scheme: dark)');
    const onMqChange = () => syncTheme();
    mq.addEventListener('change', onMqChange);

    const observer = new MutationObserver(() => syncTheme());
    observer.observe(document.documentElement, { attributes: true, attributeFilter: ['data-theme'] });

    const unsub = subscribe((v) => { vm = v; });
    const onResize = () => setViewport(window.innerWidth, window.innerHeight);
    window.addEventListener('resize', onResize);
    // A key belongs to whatever the visitor is typing into: without this, the space bar never
    // reached a single input on the desktop.
    const isTyping = (target: EventTarget | null) => {
      const el = target as HTMLElement | null;
      if (!el) return false;
      const tag = el.tagName;
      return tag === 'INPUT' || tag === 'TEXTAREA' || tag === 'SELECT' || el.isContentEditable;
    };

    const onKey = (e: KeyboardEvent) => {
      if (isTyping(e.target)) return;
      if (e.key === 'Escape') {
        os()?.closeFocused?.();
        return;
      }
      if (e.key === ' ') {
        // Spacebar: toggle ephemeral preview on selected finder file
        e.preventDefault();
        const focused = vm?.windows?.find((w) => w.id === vm?.focusedId);
        if (focused?.content?.kind === 'finder') {
          const sel = focused.content.selectedEntry;
          if (sel) {
            // Check if ephemeral already open — toggle close
            const ephemeralWin = vm?.windows?.find((w) => w.ephemeral);
            if (ephemeralWin) {
              os()?.ephemeralClose?.();
            } else {
              os()?.ephemeralPreview?.(sel);
            }
          }
        } else if (focused?.ephemeral) {
          // Spacebar on ephemeral window itself — close it
          os()?.ephemeralClose?.();
        }
      }
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
  <MenuBar menuBar={vm.menuBar} dark={isDarkMode} />
  <DesktopIcons icons={vm.desktop.icons} />
  <WindowLayer windows={vm.windows} focusedId={vm.focusedId} />
  {#if vm?.windows?.some((w) => w.ephemeral)}
    <div
      class="ephemeral-backdrop"
      onclick={() => os()?.ephemeralClose?.()}
      onkeydown={(e) => { if (e.key === ' ') { e.preventDefault(); os()?.ephemeralClose?.(); } }}
      role="button"
      tabindex={-1}
      aria-label="Close preview"
    ></div>
  {/if}
  <Dock dock={vm.dock} />
{/if}
