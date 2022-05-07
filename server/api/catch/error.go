package catch

import (
	"MeshNote/server/fos"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequestError(GinContext *gin.Context, err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
		GinContext.String(http.StatusBadRequest, err.Error())
	}
}

func HandleServerError(GinContext *gin.Context, err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
		GinContext.String(http.StatusInternalServerError, err.Error())
	}
}
