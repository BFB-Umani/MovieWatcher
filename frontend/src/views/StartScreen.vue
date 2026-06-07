<script lang="ts" setup>
import ContentCard from '../components/ContentCard.vue';
import movieTime from '../assets/images/Movie-Time.jpg'
import tvTime from '../assets/images/tv-time.jpg'

import { useRouter } from 'vue-router'

const router = useRouter()

function goToMovieListView() {
    router.push('/movies')
}
function goToSerieListView() {
    router.push('/series')
}
</script>

<template>
    <div>
        <h1 id="start-h1">Do you want to watch a movie or a TV show?</h1>
        <div id="choice-container">
            <ContentCard :-poster-path="movieTime" -title="Movie Time!" @on-click-card="goToMovieListView"/>
            <ContentCard :-poster-path="tvTime" -title="TV Time!" @on-click-card="goToSerieListView"/>
        </div>
    </div>
</template>

<style>
#choice-container {
    display: flex;
    justify-content: space-around;
    margin-top: 15%;
}
#start-h1 {
    padding-top: 40px;
    margin: 0px;
}
</style>