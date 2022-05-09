package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"MeshNote/server/fos"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Write(c *gin.Context) {
	targetId, err := para.GetInt("id", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetText, err := para.GetString("text", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetType, err := db.GetElementType(targetId)
	if targetType != db.TYPE_FILE {
		err := errors.New("read: target is not a readable file")
		catch.HandleRequestError(c, err)
		return
	} else if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	targetFileName, err := db.GetElementFileName(targetId)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	if err := fos.WriteFile(targetFileName, &targetText); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	if err := db.UpdateWordCount(targetId, fos.WordCount(&targetText)); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, "Successful!")
}
