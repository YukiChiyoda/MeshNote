package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	//path := "./data/" + c.PostForm("FileName")
	id, _ := strconv.Atoi(c.PostForm("id"))
	path, err := GetFileName(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		log.Panic(err)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	buffer := make([]byte, 100)
	for {
		reader, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				c.String(http.StatusBadRequest, err.Error())
				log.Panic(err)
			}
			break
		}
		c.String(http.StatusOK, string(buffer[:reader]))
	}
}
