<script lang="ts" setup>
import { ref } from "vue";

const searchTerm = ref("")

</script>

<template>
    <div id="input-container">
        <input v-model="searchTerm" name="search" placeholder="Search.." @keyup.enter="$emit('onSearch', searchTerm)">
        <div id="search-button" type="button">
            <img id="search-icon" src="../assets/icons/search-alt.svg" @click="$emit('onSearch', searchTerm)">
        </div>
    </div>
</template>

<style>
#input-container {
    display:flex;
    justify-content: center;
}
#search-button {
    position: relative;
    background-color: rgb(56, 77, 112);
    left: -60px;
    border-top-right-radius: 20px;
    border-bottom-right-radius: 20px;
    border: 2px solid rgb(46, 64, 92);
    min-width: 60px;
}
#search-button:hover {
  background-color: rgb(41, 57, 82);
  cursor: pointer;
}
#search-icon {
    position: relative;
    top: 20%;
}
input {
    width: 80%;
    height: 40px;
    font-size: large;
    border-radius: 20px;
    padding-left: 10px;
    margin-left:60px;
    border: 2px solid rgb(46, 64, 92);
}
</style>