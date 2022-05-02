package api

import (
	"bufio"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Write(c *gin.Context) {
	path := "./data/" + c.PostForm("FileName")
	text := c.PostForm("Text")
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		c.String(http.StatusBadRequest, "File is Not Exist")
		return
	}
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	buffer.Write([]byte(text))
	if err := buffer.Flush(); err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	count := WordCount(&text)
	c.String(http.StatusOK, strconv.Itoa(count))
	UpdateCount(c.PostForm("FileName"), count)
}
