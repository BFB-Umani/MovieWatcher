<script lang="ts" setup>


import { computed, onMounted, ref } from "vue";
import { GetMovieDetails } from "../../wailsjs/go/main/App.js";
import BackButton from "../components/BackButton.vue";
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const movie = ref({
    ID: 0,
    Title: "",
    ReleaseDate: "",
    Overview: "",
    PosterPath: "",
    Rating: 0
})

async function fetchDetails() {
    const movieIDStr = route.params.id as string
    const movieId = Number(movieIDStr)
    const movieDetails = await GetMovieDetails(movieId) as Movie
    movie.value = {
        ID: movieDetails.ID,
        Title: movieDetails.Title,
        ReleaseDate: movieDetails.ReleaseDate,
        Overview: movieDetails.Overview,
        PosterPath: movieDetails.PosterPath,
        Rating: movieDetails.Rating
    }
}

const releaseYear = computed(() => {
    const yearMonthDay = movie.value.ReleaseDate.split("-")
    return yearMonthDay[0]
})

const roundRating = computed(() => {
    const rounded = Math.round(movie.value.Rating * 10) / 10
    return rounded.toFixed(1)
})

function gotoStreamView() {
    router.push({ path: `/stream/${movie.value.ID}` })
}

onMounted(() => {
    fetchDetails()
});

</script>

<template>
    <BackButton />
    <h1 id="header">{{ movie.Title }}</h1>

    <div id="details-container-wrap">
        <div id="details-container">
            <img id="poster-details" :src="movie.PosterPath" width="300px">
            <div id="year-title-container">
                <p id="title">{{ movie.Title }}</p>
                <div id="test">
                    <p>{{ releaseYear }}</p>
                    <p id="rating">
                        {{ roundRating }}
                        <img id="icon" src="../assets/icons/star.svg">
                    </p>
                </div>
            </div>
        </div>
        <div id="overview-container">
            <p>
                {{ movie.Overview }}
            </p>
            <div id="watch-button" @click="gotoStreamView()">Watch Movie!</div>
        </div>
    </div>
</template>

<style>
#details-container-wrap {
    display: flex;
    flex-wrap: nowrap;
    margin: 30px;
}
#details-container {
    display: flex;
    flex-direction: column;
    border: 2px solid rgb(46, 64, 92);
    box-shadow: 0 0 2rem black;
    width: 300px;
}
#overview-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border: 2px solid rgb(46, 64, 92);
    box-shadow: 0 0 2rem black;
    width:65%;
    font-size: x-large;
    padding: 30px;
}
#watch-button {
    position: relative;
    background-color: rgb(56, 77, 112);
    border-radius: 20px;
    border: 2px solid rgb(46, 64, 92);
    min-width: 60px;
    font-size: xx-large;
}
#watch-button:hover {
  background-color: rgb(41, 57, 82);
  cursor: pointer;
}
</style>