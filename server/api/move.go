package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Move(c *gin.Context) {
	from, _ := strconv.Atoi(c.PostForm("from"))
	to, _ := strconv.Atoi(c.PostForm("to"))
	flag, err := IsFile(to)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if flag && to != -1 {
		c.String(http.StatusBadRequest, "Target is not a dir!")
		return
	}
	if err := MoveFile(from, to); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "0")
}
