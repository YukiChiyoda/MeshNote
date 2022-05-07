package api

import (
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"MeshNote/server/fos"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var data db.Tree
	if temp, err := para.GetString("name", c); err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		data.Name = temp
	}
	if temp, err := para.GetInt("type", c); err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		data.Type = temp
	}
	if temp, err := para.GetInt("parent", c); err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		data.Parent = temp
	}
	if temp, err := db.GetElementType(data.Parent); err != nil {
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	} else if temp != db.Type_Dir {
		err = errors.New("create: parent is not a directory")
		handleError(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	timeStemp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	data.Uptime = timeStemp
	if data.Type == db.Type_File || data.Type == db.Type_Recyled_File {
		data.FileName = timeStemp + ".md"
	} else {
		data.FileName = db.FileName_Dir
	}
	data.Id = db.Id_Undefine
	data.FileSize = db.FileSize_Empty
	if err := db.CreateElement(data); err != nil {
		handleError(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if data.Type == db.Type_File || data.Type == db.Type_Recyled_File {
		if err := fos.WriteFile(data.FileName, nil); err != nil {
			handleError(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.String(http.StatusOK, "Successful!")
}
