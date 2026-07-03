<script lang="ts">
  import type { SplashContentVM } from '$lib/os/types';

  let { content }: { content: SplashContentVM } = $props();
</script>

<div class="splash-app">
  <img src={content.icon} alt={content.name} class="splash-icon" />
  <h2>{content.name}</h2>
  <p>Coming soon</p>
</div>
