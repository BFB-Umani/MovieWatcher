<script lang="ts" setup>


import { onMounted, ref } from "vue";
import { GetShowDetails } from "../../wailsjs/go/main/App.js";
import BackButton from "../components/BackButton.vue";
import { useRoute } from 'vue-router'

const route = useRoute()

const show = ref({
    ID: 0,
    Title: "",
    SeasonNr: "",
    EpisodeNr: 0,
    NrOfEpisodes: 0,
})

async function fetchDetails() {
    const showIDStr = route.params.id as string
    const showID = Number(showIDStr)
    const seasonNr = route.params.season as string
    const episodeNrStr = route.params.episode as string
    const episodeNr = Number(episodeNrStr)
    const nrOfEpisodesStr = route.params.episodeAmt as string
    const nrOfEpisodes = Number(nrOfEpisodesStr)
    const showDetails = await GetShowDetails(showID) as Show
    show.value = {
        ID: showDetails.ID,
        Title: showDetails.Name,
        SeasonNr: seasonNr,
        EpisodeNr: episodeNr,
        NrOfEpisodes: nrOfEpisodes,
    }
}

function getShowID() {
    return `https://vsembed.ru/embed/tv?tmdb=${show.value.ID}&season=${show.value.SeasonNr}&episode=${show.value.EpisodeNr}`
}

function isFirstEpisode() {
    return show.value.EpisodeNr < 2
}

function isLastEpisode() {
    return show.value.EpisodeNr == show.value.NrOfEpisodes
}

function gotoPreviousEpisode() {
    const currEpisode = show.value.EpisodeNr
    return show.value.EpisodeNr = currEpisode - 1
}

function gotoNextEpisode() {
    const currEpisode = show.value.EpisodeNr
    return show.value.EpisodeNr = currEpisode + 1
}

function clickBackButton() {
    const currEpisode = show.value.EpisodeNr
    return show.value.EpisodeNr = currEpisode + 1
}

onMounted(() => {
    fetchDetails()
});

</script>

<template>
    <div class="container">
        <BackButton @on-click-card="clickBackButton"/>
        <div class="buttons-and-player">
            <div class="buttons-container">
                <div v-if="!isFirstEpisode()" id="watch-episode-button" @click="gotoPreviousEpisode()">
                    < Previous 
                </div>
                <div v-if="!isLastEpisode()" id="watch-episode-button" @click="gotoNextEpisode()">
                            Next > 
                </div>
            </div>
            <div id="stream-window">
                 <iframe id="iframe" type=”text/html” width="80%" height="80%" :src=getShowID() frameborder="0"
                    allowfullscreen :key="getShowID()"/>
            </div>
        </div>
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

.container {
    display: flex;
    flex-direction: column;
}

.buttons-container {
    display: flex;
    justify-content: space-around;
    margin-bottom: 20px;
}

.buttons-and-player {
    display: flex;
    flex-direction: column;
    margin-top: 50px;
}
</style>