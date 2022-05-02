package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Data []Tree `json:"Data"`
}

func Query(c *gin.Context) {
	var data []Tree
	if c.Query("p") != "" {
		p, _ := strconv.Atoi(c.Query("p"))
		QueryFile(p, &data)
	} else {
		QueryAllFile(&data)
	}
	var text Json
	text.Data = data
	str, err := json.Marshal(text)
	if err != nil {
		log.Panic(err)
	}
	c.String(http.StatusOK, string(str))
}
