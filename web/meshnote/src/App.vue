<script setup>

import { computed, onBeforeMount, onBeforeUpdate, onMounted, ref, watch, watchEffect } from 'vue';
import axios from 'axios';

const editorText = ref("");
// watch(editorText, (newValue, oldValue) => {
//     window.alert("Submit: " + editorText.value);
// });
function saveChange(){
    window.alert("Submit: " + editorText.value);
}

const hostPath = "http://localhost:3001";

const treeData = { Data: [] };
const getTreeData = onMounted(() => {
    // How can I do this automatically when opened and saved?
    axios
        .post(hostPath + "/query.go", {
            parent: 0
        }, {
            headers: {
                "Content-type": "multipart/form-data"
            }
        })
        .then(res => (treeData.Data = JSON.parse(JSON.stringify(res.data.Data))))
        .catch(err => (console.error(err)));
});

const onShowIndex = ref(null);

const changeFile = function(i){
    if(i.Type){
        console.warn("Bad Type");
        return;
    }
    axios
        .post(hostPath + "/read.go", {
            id: i.Id
        }, {
            headers: {
                "Content-type": "multipart/form-data"
            }
        })
        .then(res => (editorText.value = res.data))
        .catch(err => (console.error(err)));
    onShowIndex.value = i;
};

const saveFile = function(){
    if(onShowIndex.value.Type){
        console.warn("Bad Type");
        return;
    }
    axios
        .post(hostPath + "/write.go", {
            id: onShowIndex.value.Id,
            text: editorText.value
        }, {
            headers: {
                "Content-type": "multipart/form-data"
            }
        })
        .then(res => (console.info(res.data)))
        .catch(err => (console.error(err)));
};

// Create Delete Move User...
// I'm Waiting for You, YukiChiyoda! -- 2022.07.06 21:00

watchEffect(() => {
    console.info(treeData);
    console.info(onShowIndex.value);
});

</script>

<template>

    <div class="app">
        <div class="header"></div>
        <div class="main">
            <div id="tree" class="tree">
                <div class="card" v-for="(item, index) in treeData.Data" :key="index">
                    <!-- How to filter [Type < 0] items? -->
                    <img class="icon" :src="'/item/' + item.Type + '.ico'">
                    <div class="title" @click="changeFile(item)">{{ item.Name }}</div>
                    <div class="counter" v-if="!item.Type">{{ item.FileSize }}</div>
                </div>
            </div>
            <div class="bar">
                <div class="top"></div>
                <textarea id="editor" class="editor" wrap="physical" v-model="editorText"></textarea>
                <div class="bottom">
                    <button @click="getTreeData()">Refresh</button>
                    <button @click="saveFile()">Save</button>
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
    width: 100%;
    border-radius: 20px;
    background-color: #4c1f24;
}

.main .tree .card .icon{
    margin-left: 10px;
    margin-right: 10px;
    max-width: 20px;
}

.main .tree .card .title{
    flex: 1;
    font-size: 1.5rem;
    font-weight: 600;
    color: #eea2a4;
    text-decoration: none;
    cursor: pointer;
}

.main .tree .card .counter{
    margin-right: 10px;
    padding: 5px;
    font-size: 1.5rem;
    color: #eea2a4;
    background-color: #4c1f24;
    border-radius: 10px;
    box-shadow: 2px 2px 5px #eea2a4,
                -1px -1px 10px #eea2a4;
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

::-webkit-scrollbar{
    width: 10px;
    height: 10px;
    background-color: #a7535a;
    border-radius: 10px;
}
::-webkit-scrollbar-thumb{
    background-color: #eea2a4;
    border-radius: 10px;
}

</style>