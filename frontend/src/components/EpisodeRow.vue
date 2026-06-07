<script lang="ts" setup>

const props = defineProps({
    Title: { type: String, required: true },
    EpisodeNr: { type: String, required: true },
    PosterPath: { type: String, required: true },
    Overview: { type: String, required: true },
})

</script>

<template>
    <div id="episode-container-wrap">
        <div id="episode-details-container">
            <img id="poster-details" :src="PosterPath" width="250px">
        </div>
        <div id="title-container">
            <p id="title">{{EpisodeNr + ". " + Title }}</p>
            <p id="overview">{{ Overview }}</p>
        </div>
    </div>
</template>

<style>
#episode-container-wrap {
    display: flex;
    flex-wrap: nowrap;
    margin: 10px;
    border-bottom-right-radius: 24px;
}

#episode-container-wrap:hover {
    box-shadow: 0 0 1rem white;
    cursor: pointer;
}

#title-container {
    display: flex;
    background-color: rgba(27, 38, 54, 1);
    flex-direction: column;
    border: 2px solid rgb(46, 64, 92);
    font-size: medium;
    border-bottom-right-radius: 24px;
    text-align: start;
    padding: 10px;
    width: 100%;
}

#overview {
    margin-top: 30px;
}
</style>