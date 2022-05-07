package api

import (
	"MeshNote/server/api/catch"
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
		catch.HandleRequestError(c, err)
		return
	} else {
		data.Name = temp
	}
	if temp, err := para.GetInt("type", c); err != nil {
		catch.HandleRequestError(c, err)
		return
	} else {
		data.Type = temp
	}
	if temp, err := para.GetInt("parent", c); err != nil {
		catch.HandleRequestError(c, err)
		return
	} else {
		data.Parent = temp
	}
	if temp, err := db.GetElementType(data.Parent); err != nil {
		catch.HandleRequestError(c, err)
		return
	} else if temp != db.TYPE_DIR {
		err = errors.New("create: parent is not a directory")
		catch.HandleRequestError(c, err)
		return
	}
	timeStemp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	data.Uptime = timeStemp
	if data.Type == db.TYPE_FILE || data.Type == db.TYPE_RECYLED_FILE {
		data.FileName = timeStemp + ".md"
	} else {
		data.FileName = db.FILENAME_DIR
	}
	data.Id = db.ID_UNDEFINE
	data.FileSize = db.FILESIZE_EMPTY
	if err := db.CreateElement(data); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	if data.Type == db.TYPE_FILE || data.Type == db.TYPE_RECYLED_FILE {
		if err := fos.WriteFile(data.FileName, nil); err != nil {
			catch.HandleServerError(c, err)
			return
		}
	}
	c.String(http.StatusOK, "Successful!")
}
