package catch

import (
	"MeshNote/server/fos"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequestError(ginContext *gin.Context, err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
		ginContext.String(http.StatusBadRequest, err.Error())
	}
}

func HandleServerError(ginContext *gin.Context, err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
		ginContext.String(http.StatusInternalServerError, err.Error())
	}
}
