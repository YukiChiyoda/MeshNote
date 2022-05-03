package api

import (
	"encoding/json"
	"log"
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
	//temp.FileName = c.PostForm("FileName")
	temp.FileName = strconv.FormatInt(time.Now().UnixMicro(), 10) + ".md"
	temp.FileSize, _ = strconv.Atoi(c.PostForm("FileSize"))
	temp.Parent, _ = strconv.Atoi(c.PostForm("Parent"))
	temp.Uptime = strconv.FormatInt(time.Now().UnixMicro(), 10)
	path := "./data/" + temp.FileName
	_, err := os.Stat(path)
	if err == nil && !os.IsNotExist(err) {
		c.String(http.StatusBadRequest, "File Has Already Existed")
		return
	}
	if temp.FileSize != -1 {
		if _, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0666); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}
	if err := CreateFile(temp); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var text Json
	text.Data = append(text.Data, temp)
	str, err := json.Marshal(text)
	if err != nil {
		log.Panic(err)
	}
	c.String(http.StatusOK, string(str))
}
