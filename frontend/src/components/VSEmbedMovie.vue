<script lang="ts" setup>

function getMovieID() {
    return "https://vsembed.ru/embed/movie?tmdb=687163"
}

</script>

<template>
    <iframe id=”player” type=”text/html” width="80%" height="80%" :src=getMovieID() frameborder="0" allowfullscreen="true" />
</template>