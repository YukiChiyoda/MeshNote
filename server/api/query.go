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
	p, err := para.GetInt("p", c)
	if err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// if p == 0 { db.QueryAllFile() }
	var t QueryJson
	t.Data, err = db.QueryFile(p)
	if err != nil {
		handleError(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	s, err := json.Marshal(t)
	if err != nil {
		handleError(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, string(s))
}
