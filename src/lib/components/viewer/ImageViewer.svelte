<script lang="ts">
  import type { ViewerContentVM } from '$lib/os/types';

  let { content }: { content: ViewerContentVM } = $props();
</script>

<div class="image-viewer">
  <img src={content.src} alt={content.name} />
</div>
