package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Drop(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	var path string
	if flag, _ := IsFile(id); flag {
		var err error
		path, err = GetFileName(id)
		if err != nil {
			log.Panic(err)
		}
	}
	if err := DropFile(id); err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	if path != "" {
		if err := os.Remove(path); err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
	c.String(http.StatusOK, "0")
}
