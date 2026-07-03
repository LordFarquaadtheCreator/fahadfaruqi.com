<script lang="ts">
  import type { ViewerContentVM } from '$lib/os/types';

  let { content }: { content: ViewerContentVM } = $props();
</script>

<div class="text-viewer-app">
  <div class="viewer-toolbar">
    <div class="viewer-toolbar-spacer"></div>
    <span class="viewer-title">{content.name}</span>
    <div class="viewer-toolbar-right">
      <button class="viewer-btn" aria-label="Share">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.3">
          <path d="M7 1V9M4 4L7 1L10 4" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M2 8V12C2 12.5 2.5 13 3 13H11C11.5 13 12 12.5 12 12V8" stroke-linecap="round"/>
        </svg>
      </button>
    </div>
  </div>
  <div class="text-viewer-content">
    {#if content.isMarkdown}
      <div class="markdown-body">{@html content.content ?? ''}</div>
    {:else}
      <pre class="plain-text">{content.content ?? ''}</pre>
    {/if}
  </div>
</div>
