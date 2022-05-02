package api

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var temp Tree
	temp.Id = -1
	temp.Name = c.PostForm("Name")
	temp.FileName = c.PostForm("FileName")
	temp.FileSize, _ = strconv.Atoi(c.PostForm("FileSize"))
	temp.Parent, _ = strconv.Atoi(c.PostForm("Parent"))
	temp.Uptime = strconv.FormatInt(time.Now().UnixMicro(), 10)
	CreateFile(temp)
	path := "./data/" + c.PostForm("FileName")
	_, err := os.Stat(path)
	if err == nil && !os.IsNotExist(err) {
		c.String(http.StatusBadRequest, "File Has Already Existed")
		return
	}
	if _, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0666); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, temp.FileName)
}
