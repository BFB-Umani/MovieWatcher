<script lang="ts" setup>


import { onMounted, ref } from "vue";
import { GetMovieDetails } from "../../wailsjs/go/main/App.js";
import BackButton from "../components/BackButton.vue";
import { useRoute } from 'vue-router'

const route = useRoute()

const movie = ref({
    ID: 0,
    Title: "",
})

async function fetchDetails() {
    const movieIDStr = route.params.id as string
    const movieId = Number(movieIDStr)
    const movieDetails = await GetMovieDetails(movieId) as Movie
    movie.value = {
        ID: movieDetails.ID,
        Title: movieDetails.Title,
    }
}

function getMovieID() {
    return "https://vsembed.ru/embed/movie?tmdb=" + movie.value.ID
}

onMounted(() => {
    fetchDetails()
});

</script>

<template>
    <BackButton />

    <div id="stream-window">
        <iframe type=”text/html” width="80%" height="80%" :src=getMovieID() frameborder="0" allowfullscreen />
    </div>

</template>

<style>
    #stream-window {
        width: 100%;
        min-height: 650px;
        height: 1000px;
        margin-bottom: 0px;
        padding-bottom: 0px;
    }
</style>