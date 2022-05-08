package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Move(c *gin.Context) {
	originElement, err := para.GetInt("from", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetParent, err := para.GetInt("to", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetType, err := db.GetElementType(targetParent)
	if targetType == db.SPECIAL_INVALID_ID {
		catch.HandleRequestError(c, err)
		return
	} else if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	if targetType != db.TYPE_DIR {
		err := errors.New("move: target is not a diretory")
		catch.HandleRequestError(c, err)
		return
	}
	if err := db.MoveElement(originElement, targetParent); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, "Successful!")
}
