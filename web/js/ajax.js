var host = "http://localhost:3000"

const TYPE_FILE = 0;
const TYPE_RECYLED_FILE = -1;
const TYPE_DIR = 1;

function getElement(parent, depth) {
    $.ajax({
        url: host + "/query.go",
        data: {
            parent: parent
        },
        dataType: "json",
        type: "POST",
        async: false,
        success: function (obj) {
            var data = obj.Data;
            if (data == null) {
                // console.error("tree bottom");
                return;
            }
            console.log(data);
            for (var i = 0; i < data.length; i++) {
                var title = "";
                for (var j = 0; j < depth; j++) {
                    title += ">";
                }
                title += "[" + data[i].Id + "]";
                title += data[i].Name;
                if (data[i].Type == TYPE_FILE) {
                    title += "<count=" + data[i].FileSize + ">";
                }
                appendCard(data[i].Id, title)
                if (data[i].Type == TYPE_DIR) {
                    getElement(data[i].Id, depth + 1)
                }
            }
            return;
        },
        error: function (data, type, err) {
            console.error(data);
            console.error(type);
            console.error(err);
        }
    });
}

function getText(id) {
    $.ajax({
        url: host + "/read.go",
        data: {
            id: id
        },
        dataType: "text",
        type: "POST",
        async: false,
        success: function (obj) {
            if (obj == null) {
                console.error("-1");
                return null;
            }
            writeText(id, obj);
            return;
        },
        error: function (data, type, err) {
            console.error(data);
            console.error(type);
            console.error(err);
        }
    });
}