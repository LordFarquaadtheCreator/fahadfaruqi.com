<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import type { MenuBarVM } from '$lib/os/types';

  let { menuBar }: { menuBar: MenuBarVM } = $props();

  let openMenu = $state<number | null>(null);
  let clock = $state('');

  function updateClock() {
    const now = new Date();
    const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
    const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
    let h = now.getHours();
    const ampm = h >= 12 ? 'PM' : 'AM';
    h = h % 12 || 12;
    const m = now.getMinutes().toString().padStart(2, '0');
    clock = `${days[now.getDay()]} ${months[now.getMonth()]} ${now.getDate()} ${h}:${m} ${ampm}`;
  }

  let timer: ReturnType<typeof setInterval>;

  onMount(() => {
    updateClock();
    timer = setInterval(updateClock, 1000);
  });

  onDestroy(() => clearInterval(timer));

  function toggleMenu(i: number) {
    openMenu = openMenu === i ? null : i;
  }
</script>

<div class="menubar">
  <div class="menubar-left">
    <button class="menu-item apple-logo" class:active={openMenu === 0} onclick={() => toggleMenu(0)} aria-label="Apple menu">
      <svg width="14" height="16" viewBox="0 0 14 16" fill="currentColor">
        <path d="M11.2 8.4c0-1.9 1.6-2.8 1.7-2.8-0.9-1.3-2.3-1.5-2.8-1.5-1.2-0.1-2.3 0.7-2.9 0.7-0.6 0-1.5-0.7-2.5-0.7-1.3 0-2.5 0.7-3.1 1.9-1.3 2.3-0.3 5.7 1 7.5 0.6 0.9 1.3 1.9 2.3 1.9 0.9 0 1.3-0.6 2.4-0.6 1.1 0 1.4 0.6 2.4 0.6 1 0 1.6-0.9 2.2-1.8 0.7-1 1-2 1-2.1-0.1 0-1.7-0.7-1.7-2.6zM9.5 2.8c0.5-0.6 0.9-1.5 0.8-2.4-0.8 0-1.7 0.5-2.2 1.2-0.5 0.5-0.9 1.4-0.8 2.3 0.9 0.1 1.7-0.4 2.2-1.1z"/>
      </svg>
    </button>
    {#each menuBar.menus.slice(1) as menu, i}
      <button
        class="menu-item"
        class:app-name={i === 0}
        class:active={openMenu === i + 1}
        onclick={() => toggleMenu(i + 1)}
      >
        {menu.name}
      </button>
    {/each}
  </div>
  <div class="menubar-right">
    {#each menuBar.tray as item}
      <span class="tray-icon" data-icon={item.icon}>
        <span class="tray-tooltip">{item.name}</span>
        {#if item.icon === 'search'}
          <svg width="15" height="15" viewBox="0 0 15 15" fill="none" stroke="currentColor" stroke-width="1.5">
            <circle cx="6" cy="6" r="4.5"/>
            <line x1="9.5" y1="9.5" x2="13" y2="13" stroke-linecap="round"/>
          </svg>
        {:else if item.icon === 'controls'}
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none" stroke="currentColor" stroke-width="1.3">
            <circle cx="5" cy="5" r="2.5"/>
            <line x1="7" y1="5" x2="13" y2="5" stroke-linecap="round"/>
            <circle cx="11" cy="11" r="2.5"/>
            <line x1="3" y1="11" x2="9" y2="11" stroke-linecap="round"/>
          </svg>
        {:else if item.icon === 'wifi'}
          <svg width="17" height="15" viewBox="0 0 17 15" fill="currentColor">
            <path d="M8.5 1C5.2 1 2.2 2.2 0 4.2L1.5 5.7C3.3 4.1 5.8 3.1 8.5 3.1s5.2 1 7 2.6L17 4.2C14.8 2.2 11.8 1 8.5 1z"/>
            <path d="M8.5 4.6C6.2 4.6 4.1 5.5 2.6 7L4.1 8.5C5.2 7.5 6.8 6.9 8.5 6.9s3.3 0.6 4.4 1.6L14.4 7C12.9 5.5 10.8 4.6 8.5 4.6z"/>
            <path d="M8.5 8.2C7 8.2 5.7 8.8 4.8 9.8L8.5 13.5L12.2 9.8C11.3 8.8 10 8.2 8.5 8.2z"/>
          </svg>
        {:else if item.icon === 'battery-full'}
          <svg width="22" height="12" viewBox="0 0 22 12" fill="currentColor">
            <rect x="0.5" y="1" width="18" height="10" rx="2" fill="none" stroke="currentColor" stroke-width="1"/>
            <rect x="2" y="2.5" width="15" height="7" rx="1"/>
            <rect x="19.5" y="3.5" width="2" height="5" rx="1"/>
          </svg>
        {/if}
      </span>
    {/each}
    <span class="clock">{clock}</span>
  </div>
</div>
