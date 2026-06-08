<script lang="ts" setup>
import { computed } from 'vue';

const props = defineProps({
    ID: { type: Number, required: false },
    Title: { type: String, required: false },
    Name: { type: String, required: false },
    ReleaseDate: { type: String, required: false },
    PosterPath: { type: String, required: true },
    Rating: {type: Number, required: false}
})

const releaseYear = computed(() => {
    if (props.ReleaseDate) {
        const yearMonthDay = props.ReleaseDate.split("-")
        return yearMonthDay[0]
    }
    return ""
})

const getTitle = computed(() => {
    return props.Title ? props.Title : props.Name
})

const roundRating = computed(() => {
    if (props.Rating) {
        const rounded = Math.round(props.Rating * 10) / 10
        return rounded.toFixed(1)
    }
    return ""
})

const imageWidth = computed(() => {
    if (props.ID) {
        return "300px"
    }
    return "450px"
})

const imageWContainerWidth = computed(() => {
    if (props.ID) {
        return "304px"
    }
    return "454px"
})

</script>

<template>
    <div id="card-container" @click="$emit('onClickCard', ID)">
        <img id="poster" :src="PosterPath" :width="imageWidth">
        <div id="year-title-container">
            <p id="title">{{ getTitle }}</p>
            <div id="year-and-rating">
                <p>{{ releaseYear }}</p>
                <p v-if="Rating" id="rating">
                    {{ roundRating }}
                    <img id="icon" src="../assets/icons/star.svg">
                </p>

            </div>
        </div>
    </div>
</template>

<style>
#card-container {
    display: flex;
    flex-direction: column;
    border-radius: 25px;
    border: 2px solid rgb(46, 64, 92);
    margin: 30px;
    box-shadow: 0 0 2rem black;
    width: v-bind('imageWContainerWidth')
}

#card-container:hover {
    box-shadow: 0 0 1rem white;
    cursor: pointer;
}

#poster {
    border-top-left-radius: 24px;
    border-top-right-radius: 24px;
}

#year-title-container {
    display: flex;
    background-color: rgba(27, 38, 54, 1);
    flex-direction: column;
    border-top: 2px solid rgb(46, 64, 92);
    font-size: large;
    border-bottom-left-radius: 24px;
    border-bottom-right-radius: 24px;
}

#title {
    text-wrap: wrap;
    font-size: x-large;
    font-weight: bold;
    margin-top: 10px;
    margin-bottom: -20px;
    width: 100%
}
#year-and-rating {
    display: flex;
    justify-content: space-around;
    align-items: center;
}
#rating {
    position: relative;
    top: -3px;
}
#icon {
    position: relative;
    top: 5px;
}
</style>