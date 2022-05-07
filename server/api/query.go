package api

import (
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"MeshNote/server/fos"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryJson struct {
	Data []db.Tree `json:"Data"`
}

func handleError(err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
	}
}

func Query(c *gin.Context) {
	parent, err := para.GetInt("parent", c)
	if err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// log.Println(parent)
	// if parent == 0 { db.QueryAllFile() }
	var jsonData QueryJson
	jsonData.Data, err = db.QueryElement(parent)
	if err != nil {
		handleError(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	jsonText, err := json.Marshal(jsonData)
	if err != nil {
		handleError(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, string(jsonText))
}
