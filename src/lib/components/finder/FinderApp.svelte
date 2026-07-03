<script lang="ts">
  import type { WindowVM, FinderContentVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { window: win, content }: { window: WindowVM; content: FinderContentVM } = $props();

  function onSidebarClick(path: string) {
    os()?.finderNavigate?.(win.id, path);
  }

  function onEntryClick(path: string, isDir: boolean) {
    os()?.finderSelect?.(win.id, path);
  }

  function onEntryDoubleClick(path: string) {
    os()?.finderOpenEntry?.(win.id, path);
  }

  function onBack() {
    os()?.finderBack?.(win.id);
  }

  function onForward() {
    os()?.finderForward?.(win.id);
  }
</script>

<div class="finder-app">
  <div class="finder-sidebar">
    {#each content.sidebar as section}
      <div class="sidebar-section">
        {#if section.title}
          <div class="sidebar-section-title">{section.title}</div>
        {/if}
        {#each section.items as item}
          <button
            class="sidebar-item"
            class:active={item.active}
            onclick={() => onSidebarClick(item.path)}
          >
            <img class="sidebar-item-icon" src={item.icon} alt="" />
            <span>{item.name}</span>
          </button>
        {/each}
      </div>
    {/each}
  </div>
  <div class="finder-main">
    <div class="finder-toolbar">
      <button class="nav-btn" disabled={!content.canBack} onclick={onBack} aria-label="Back">
        <svg width="12" height="12" viewBox="0 0 12 12" fill="none" stroke="currentColor" stroke-width="1.8">
          <path d="M8 2L3 6L8 10" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </button>
      <button class="nav-btn" disabled={!content.canForward} onclick={onForward} aria-label="Forward">
        <svg width="12" height="12" viewBox="0 0 12 12" fill="none" stroke="currentColor" stroke-width="1.8">
          <path d="M4 2L9 6L4 10" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </button>
      <div class="toolbar-spacer"></div>
      <button class="nav-btn" aria-label="View options">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor">
          <rect x="1" y="2" width="12" height="2" rx="1"/>
          <rect x="1" y="6" width="12" height="2" rx="1"/>
          <rect x="1" y="10" width="12" height="2" rx="1"/>
        </svg>
      </button>
      <button class="nav-btn" aria-label="Search">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.5">
          <circle cx="6" cy="6" r="4"/>
          <line x1="9" y1="9" x2="12" y2="12" stroke-linecap="round"/>
        </svg>
      </button>
    </div>
    {#if content.content.kind === 'folder'}
      <div class="finder-listing">
        {#each content.content.entries ?? [] as entry}
          <button
            class="entry"
            class:selected={entry.path === content.selectedEntry}
            onclick={() => onEntryClick(entry.path, entry.kind === 'dir')}
            ondblclick={() => onEntryDoubleClick(entry.path)}
          >
            <img class="entry-icon" src={entry.icon} alt="" />
            <span class="entry-name">{entry.name}</span>
          </button>
        {/each}
      </div>
    {:else if content.content.kind === 'markdown'}
      <div class="finder-preview">{@html content.content.html ?? ''}</div>
    {/if}
    <div class="finder-breadcrumb">
      {#each content.breadcrumb as crumb, i}
        {#if i > 0}
          <span class="breadcrumb-sep">›</span>
        {/if}
        <button class="breadcrumb-item" onclick={() => onSidebarClick(crumb.path)}>
          <img class="breadcrumb-icon" src={crumb.icon} alt="" />
          <span>{crumb.name}</span>
        </button>
      {/each}
    </div>
  </div>
</div>
