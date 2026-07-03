<script lang="ts">
  let { src }: { src: string } = $props();
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
