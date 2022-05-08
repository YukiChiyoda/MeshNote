package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryJson struct {
	Data []db.Tree `json:"Data"`
}

func Query(c *gin.Context) {
	target, err := para.GetInt("parent", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	// log.Println(parent)
	// if parent == 0 { db.QueryAllFile() }
	var jsonData QueryJson
	jsonData.Data, err = db.QueryElement(target)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	jsonText, err := json.Marshal(jsonData)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, string(jsonText))
}
