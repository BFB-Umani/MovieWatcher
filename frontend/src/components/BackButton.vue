<template>
    <div id="back-container">
        <img id="back" src="../assets/icons/arrow-back.svg" @click="$router.back">
    </div>
</template>

<style>
#back-container {
    display: flex;
    flex-wrap: wrap;
    width: 80%;
    margin-left: 20px;
    margin-bottom: -40px;
}

#back {
    width: 5%;
    border: 2px solid rgb(46, 64, 92);
    background-color: rgb(56, 77, 112);
    border-radius: 50px;
    margin-top: 20px;
}

#back:hover {
    background-color: rgb(41, 57, 82);
    cursor: pointer;
}
</style>