<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from "vue";
import { GetTrendingMovies, SearchMovie } from "../../wailsjs/go/main/App.js";
import ContentCard from "../components/ContentCard.vue";
import BackButton from "../components/BackButton.vue";
import SearchBar from "../components/SearchBar.vue";
import { useRouter } from 'vue-router'

const router = useRouter()

const movies = reactive({
    list: [] as Movie[]
})

const isTrendingSearch = ref(true)

async function fetchTrending() {
    const TrendingMovies = await GetTrendingMovies() as Movie[]
    movies.list = TrendingMovies
}

function clickMovieCard(ID: string) {
    router.push({ path: `/movies/${ID}` })
}

const headerText = computed(() => {
    if (isTrendingSearch.value) {
        return "Trending movies this week"
    } else if (movies.list.length === 0) {
        return "No search results found"
    } else {
        return "Showing search results:"
    }
})

async function searchMovie(searchTerm: string) {
    if (searchTerm == "") {
        isTrendingSearch.value = true
        const searchResult = await GetTrendingMovies() as Movie[]
        movies.list = searchResult
    } else {
        isTrendingSearch.value = false
        const searchResult = await SearchMovie(searchTerm) as Movie[]
        movies.list = searchResult
    }
}

onMounted(() => {
    fetchTrending()
});

</script>

<template>
    <BackButton />
    <h1 id="header">{{ headerText }}</h1>
    <SearchBar @on-search="searchMovie" />
    <div id="container">
        <div v-for="result in movies.list">
            <ContentCard :-title="result.Title" :-i-d="result.ID" :-poster-path="result.PosterPath"
                :-release-date="result.ReleaseDate" :-rating="result.Rating" @on-click-card="clickMovieCard" />
        </div>
    </div>
</template>

<style>
#container {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
}
</style>