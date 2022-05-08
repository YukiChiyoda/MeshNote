package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"MeshNote/server/fos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id, err := para.GetInt("id", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	fileName, err := db.GetElementFileName(id)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	fileType, err := db.GetElementType(id)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	if fileType == db.TYPE_FILE {
		if err := fos.DeleteFile(fileName); err != nil {
			catch.HandleServerError(c, err)
			return
		}
	}
	if err := db.DeleteElement(id); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, "Successful!")
}
