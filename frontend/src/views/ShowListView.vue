<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from "vue";
import { GetTrendingShows, SearchShow } from "../../wailsjs/go/main/App.js";
import ContentCard from "../components/ContentCard.vue";
import BackButton from "../components/BackButton.vue";
import SearchBar from "../components/SearchBar.vue";
import { useRouter } from 'vue-router'

const router = useRouter()

const shows = reactive({
    list: [] as Show[]
})

const isTrendingSearch = ref(true)

async function fetchTrending() {
    const TrendingShows = await GetTrendingShows() as Show[]
    shows.list = TrendingShows
}

function clickShowCard(ID: string) {
    router.push({ path: `/series/${ID}` })
}

const headerText = computed(() => {
    if (isTrendingSearch.value) {
        return "Trending TV shows this week"
    } else if (shows.list.length === 0) {
        return "No search results found"
    } else {
        return "Showing search results:"
    }
})

async function searchShow(searchTerm: string) {
    if (searchTerm == "") {
        isTrendingSearch.value = true
        const searchResult = await GetTrendingShows() as Show[]
        shows.list = searchResult
    } else {
        isTrendingSearch.value = false
        const searchResult = await SearchShow(searchTerm) as Show[]
        shows.list = searchResult
    }
}

onMounted(() => {
    fetchTrending()
});

</script>

<template>
    <BackButton />
    <h1 id="header">{{ headerText }}</h1>
    <SearchBar @on-search="searchShow" />
    <div id="container">
        <div v-for="result in shows.list">
            <ContentCard :-name="result.Name" :-i-d="result.ID" :-poster-path="result.PosterPath"
                :-release-date="result.ReleaseDate" :-rating="result.Rating" @on-click-card="clickShowCard" />
        </div>
    </div>
</template>

<style>
#container {
    display: flex;
    flex-wrap: wrap;
}
</style>