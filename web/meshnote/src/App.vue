<script setup>

import { computed, onMounted, ref } from 'vue';
import axios from 'axios';

const editorText = ref("Test");
// watch(editorText, (newValue, oldValue) => {
//     window.alert("Submit: " + editorText.value);
// });
function saveChange(){
    window.alert("Submit: " + editorText.value);
}

const hostPath = "http://localhost:3001";
const getTreeData = onMounted(() => {
    axios
        .post(hostPath + "/query.go", {
            parent: 0
        })
        .then(res => (treeData = JSON.parse(res)))
        .catch(err => (console.error(err)));
})

const treeData = ref()

const selectIcon = computed(() => {
    const iconPath = "/icon/";
    return iconPath + 1 + ".ico";
})

</script>

<template>

    <div class="app">
        <div class="header"></div>
        <div class="main">
            <div id="tree" class="tree">
                <div class="card">
                    <img class="icon" :src="selectIcon">
                    <div class="title">MeshNote</div>
                    <div class="counter">0</div>
                </div>
            </div>
            <div class="bar">
                <div class="top"></div>
                <textarea id="editor" class="editor" wrap="physical" v-model="editorText"></textarea>
                <div class="bottom">
                    <button @click="saveChange">Save</button>
                </div>
            </div>
        </div>
        <div class="footer"></div>
    </div>

</template>

<style>

* {
    margin: 0;
    padding: 0;
}

textarea {
    border: none;
    outline: none;
    resize: none;
}

#vue {
    display: -webkit-flex;
    display: flex;
    justify-content: center;
    /* background-image: url(./img/pattern.png); */
    background-color: #f0c9cf;
    overflow: hidden;
}

.app {
    display: -webkit-flex;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-width: 60%;
    max-width: 80%;
    height: 100vh;
}

.header {
    flex: 0 0 100px;
    width: 100%;
    background-color: #a7535a;
}

.main {
    display: flex;
    flex: 1;
    width: 100%;
}

.main .tree {
    display: flex;
    flex: 0 0 25%;
    flex-direction: column;
    justify-content: start;
    align-items: center;
    padding: 10px;
    background-color: #eea2a4;
    overflow-x: hidden;
    overflow-y: scroll;
}

.main .tree .card {
    display: flex;
    align-items: center;
    flex: 0 0 5%;
    margin: 5px;
    padding: 2px 5px 2px 5px;   
    width: 100%;
    background-color: #4c1f24;
}

.main .tree .card .icon{
    max-width: 5%;
}

.main .tree .card .title{
    flex: 1;
    font-size: 1.5rem;
    font-weight: 600;
    color: #eea2a4;
}

.main .tree .card .counter{
    font-size: 1.5rem;
    color: #eea2a4;
}

.main .bar {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.main .bar .top {
    flex: 0 0 5%;
    background-color: #eea2a4;
}

.main .bar .editor {
    flex: 1;
    background-color: #4c1f24;
    overflow-x: hidden;
    overflow-y: scroll;
    padding: 10px;
    color: #eea2a4;
    font-size: 2rem;
}

.main .bar .bottom {
    flex: 0 0 5%;
    background-color: #eea2a4;
}

.footer {
    flex: 0 0 100px;
    width: 100%;
    background-color: #a7535a;
}

</style>