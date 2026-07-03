<script lang="ts">
  import type { EphemeralContentVM } from '$lib/os/types';
  import { os } from '$lib/os/osStore.svelte';

  let { content }: { content: EphemeralContentVM } = $props();
</script>

{#if content.kind === 'ephemeral-folder'}
  <div class="ephemeral-preview ephemeral-folder">
    <img class="ephemeral-folder-icon" src={content.icon} alt="" />
    <span class="ephemeral-folder-name">{content.name}</span>
  </div>
{:else}
  <div class="ephemeral-preview">
    <img src={content.src} alt={content.name} />
    <button
      class="ephemeral-open-preview"
      onclick={() => os()?.ephemeralOpenPreview?.()}
    >
      Open in Preview
    </button>
  </div>
{/if}
