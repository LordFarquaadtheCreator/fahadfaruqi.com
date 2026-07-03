<script lang="ts">
  import type { WindowVM } from '$lib/os/types';

  let {
    window: win,
    focused = false,
    onclose,
    onfocus,
    ondrag,
    onresize,
    onminimize,
    onmaximize,
    children,
  }: {
    window: WindowVM;
    focused?: boolean;
    onclose: () => void;
    onfocus: () => void;
    ondrag: (dx: number, dy: number) => void;
    onresize: (w: number, h: number) => void;
    onminimize: () => void;
    onmaximize: () => void;
    children: import('svelte').Snippet;
  } = $props();

  let dragging = $state(false);
  let lastX = 0;
  let lastY = 0;

  function onPointerDown(e: PointerEvent) {
    if (e.target instanceof HTMLElement && e.target.closest('.traffic-light')) return;
    if (e.target instanceof HTMLElement && e.target.closest('.window-content')) {
      onfocus();
      return;
    }
    onfocus();
    dragging = true;
    lastX = e.clientX;
    lastY = e.clientY;
    e.preventDefault();
  }

  function onPointerMove(e: PointerEvent) {
    if (!dragging) return;
    const dx = e.clientX - lastX;
    const dy = e.clientY - lastY;
    lastX = e.clientX;
    lastY = e.clientY;
    ondrag(dx, dy);
  }

  function onPointerUp() {
    dragging = false;
  }

  let resizing = $state(false);
  let resizeStartX = 0;
  let resizeStartY = 0;
  let resizeStartW = 0;
  let resizeStartH = 0;

  function onResizeStart(e: PointerEvent) {
    e.stopPropagation();
    resizing = true;
    resizeStartX = e.clientX;
    resizeStartY = e.clientY;
    resizeStartW = win.width;
    resizeStartH = win.height;
  }

  function onResizeMove(e: PointerEvent) {
    if (!resizing) return;
    const w = Math.max(300, resizeStartW + (e.clientX - resizeStartX));
    const h = Math.max(200, resizeStartH + (e.clientY - resizeStartY));
    onresize(w, h);
  }

  function onResizeEnd() {
    resizing = false;
  }

  function attachGlobalListeners() {
    window.addEventListener('pointermove', onPointerMove);
    window.addEventListener('pointerup', onPointerUp);
    window.addEventListener('pointermove', onResizeMove);
    window.addEventListener('pointerup', onResizeEnd);
  }

  function detachGlobalListeners() {
    window.removeEventListener('pointermove', onPointerMove);
    window.removeEventListener('pointerup', onPointerUp);
    window.removeEventListener('pointermove', onResizeMove);
    window.removeEventListener('pointerup', onResizeEnd);
  }

  $effect(() => {
    if (dragging || resizing) {
      attachGlobalListeners();
      return detachGlobalListeners;
    }
  });
</script>

<div
  class="window"
  class:focused
  class:ephemeral={win.ephemeral}
  style="left: {win.x}px; top: {win.y}px; width: {win.width}px; height: {win.height}px; z-index: {win.z};"
  onpointerdown={onPointerDown}
  role="group"
  aria-label={win.title}
>
  {#if !win.ephemeral}
    <div class="titlebar">
      <div class="traffic-lights">
        <button class="traffic-light close" onclick={onclose} aria-label="Close"></button>
        <button class="traffic-light minimize" onclick={onminimize} aria-label="Minimize"></button>
        <button class="traffic-light maximize" onclick={onmaximize} aria-label="Maximize"></button>
      </div>
      <span class="title">{win.title}</span>
    </div>
  {/if}
  <div class="window-content">
    {@render children()}
  </div>
  {#if !win.ephemeral}
    <div class="resize-handle" role="separator" onpointerdown={onResizeStart}></div>
  {/if}
</div>
