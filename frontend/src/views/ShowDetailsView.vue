<script lang="ts" setup>

import { computed, getCurrentInstance, onMounted, ref } from "vue";
import { GetShowDetails, GetSeasonEpisodes } from "../../wailsjs/go/main/App.js";
import BackButton from "../components/BackButton.vue";
import { useRoute, useRouter } from 'vue-router'
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';
import EpisodeRow from "../components/EpisodeRow.vue";

const router = useRouter()
const route = useRoute()

const show = ref({
    ID: 0,
    Title: "",
    ReleaseDate: "",
    Overview: "",
    PosterPath: "",
    Rating: 0,
    Seasons: [] as Season[]
})

const episodes = ref({} as { [key: string]: Episode[] })

async function fetchDetails() {
    const showIDStr = route.params.id as string
    const showId = Number(showIDStr)
    const showDetails = await GetShowDetails(showId) as Show
    const filteredSeasons = showDetails.Seasons!.filter((season) => season.name != "Specials")
    const seasonEpisodes = await GetSeasonEpisodes(showId, filteredSeasons[0].season_number) as Episode[]
    show.value = {
        ID: showDetails.ID,
        Title: showDetails.Name,
        ReleaseDate: showDetails.ReleaseDate,
        Overview: showDetails.Overview,
        PosterPath: showDetails.PosterPath,
        Rating: showDetails.Rating,
        Seasons: filteredSeasons
    }
    episodes.value[filteredSeasons[0].season_number] = seasonEpisodes
}

const releaseYear = computed(() => {
    const yearMonthDay = show.value.ReleaseDate.split("-")
    return yearMonthDay[0]
})

const roundRating = computed(() => {
    const rounded = Math.round(show.value.Rating * 10) / 10
    return rounded.toFixed(1)
})

function gotoStreamView(seasonNr: number, episodeNr: number, episodeAmt: number) {
    router.push({ path: `/series/stream/${show.value.ID}/${seasonNr}/${episodeNr}/${episodeAmt}/` })
}

async function fetchSeasonEpisodes(seasonNr: string) {
    const seasonNumber = Number(seasonNr) + 1
    if (!episodes.value[seasonNumber]) {
        const seasonEpisodes = await GetSeasonEpisodes(show.value.ID, seasonNumber)
        episodes.value[seasonNumber] = seasonEpisodes
    }
}

onMounted(() => {
    fetchDetails()
});

</script>

<template>
    <BackButton />
    <h1 id="header">{{ show.Title }}</h1>

    <div id="show-details-view">
        <div id="details-container-wrap">
            <div id="details-container">
                <img id="poster-details" :src="show.PosterPath" width="300px">
                <div id="year-title-container">
                    <p id="title">{{ show.Title }}</p>
                    <div id="test">
                        <p>{{ releaseYear }}</p>
                        <p id="rating">
                            {{ roundRating }}
                            <img id="icon" src="../assets/icons/star.svg">
                        </p>
                    </div>
                </div>
            </div>
            <Tabs :value="0" scrollable id="tabs-container"
                @update:value="(value) => fetchSeasonEpisodes(String(value))">
                <TabList>
                    <Tab v-for="season, index in show.Seasons" :key="index" :value="season.season_number - 1">{{ 'season ' + season.season_number }}</Tab>
                </TabList>
                <TabPanels id="tab-panels">
                    <TabPanel v-for="season, index in show.Seasons" :key="season.id" :value="index" id="tab-panel">
                        <p class="overview-text">{{ season.overview || "No overview found" }}</p>
                        <div id="episodes-container" ref="episodesRef">
                            <p class="episodes-text">{{ "Episodes: " }}</p>
                            <EpisodeRow v-for="episode in episodes[season.season_number]" :-title="episode.Name"
                                :-poster-path="episode.StillPath" :-overview="episode.Overview"
                                :-episode-nr="episode.EpisodeNumber"
                                @click="gotoStreamView(season.season_number, episode.EpisodeNumber, season.episode_count)">
                            </EpisodeRow>
                        </div>
                    </TabPanel>
                </TabPanels>
            </Tabs>
        </div>
    </div>
</template>

<style>
#show-details-view {
    display: flex;
    flex-direction: column;
    flex-wrap: nowrap;
    margin: 30px;
    width: 90%;
    height: 100%;
}

#details-container-wrap {
    display: flex;
    flex-wrap: nowrap;
    height: 100%;
    max-height: 900px;
}

#episodes-container {
    margin-bottom: 40px;
    overflow: scroll;
    height: 100%;
}

#tabs-container {
    width: 80%;
}

#tab-panels {
    height: 100%;
    overflow: scroll;
}

.overview-text {
    font-size: large;
}
.episodes-text {
    font-size: x-large;
    text-align: start;
    font-weight: bold;
}

#watch-episode-button {
    background-color: rgb(56, 77, 112);
    border-radius: 20px;
    border: 2px solid rgb(46, 64, 92);
    width: 100px;
    font-size: medium;
    margin: 3px;
}

#watch-episode-button:hover {
    background-color: rgb(41, 57, 82);
    cursor: pointer;
}
</style>