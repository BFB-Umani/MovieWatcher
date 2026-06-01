<script lang="ts" setup>

function getMovieID() {
  return "SERIES"
}

</script>

<template>
    <iframe id=”player” type=”text/html” width="100%" height="100%" :src=getMovieID() frameborder="0" allowfullscreen/>
</template>