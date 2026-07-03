<script lang="ts">
  // Wallpaper path served from static/wallpapers/
  let { src = '/wallpapers/wallpaper.jpg' }: { src?: string } = $props();
</script>

<div class="wallpaper" style="background-image: url({src})"></div>

<style>
  .wallpaper {
    position: fixed;
    inset: 0;
    background-size: cover;
    background-position: center;
    z-index: -1;
  }
</style>
