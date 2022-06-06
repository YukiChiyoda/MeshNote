$(function () {
    getElement(0, 0);
})

function cardClickEvent(obj) {
    getText(obj.target.id);
}

function appendCard(id, title) {
    var temp = document.createElement("div");
    temp.id = id;
    temp.className = "card";
    temp.innerText = title;
    temp.onclick = cardClickEvent;
    $("#tree").append(temp);
}

function writeText(id, text) {
    $("#editor").text(text);
}