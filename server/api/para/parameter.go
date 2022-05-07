// Examinate parameters and convert their types

package para

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInt(ParaName string, GinContext *gin.Context) (int, error) {
	t := GinContext.PostForm(ParaName)
	if t != "" {
		return strconv.Atoi(t)
	} else {
		return 0, fmt.Errorf("para: lost parameter `%s`", ParaName)
	}
}

/*
func GetIntCanBeEmpty(ParaName string, GinContext *gin.Context) (int, error) {
	return strconv.Atoi(GinContext.PostForm(ParaName))
}
*/

func GetString(ParaName string, GinContext *gin.Context) (string, error) {
	t := GinContext.PostForm(ParaName)
	if t != "" {
		return t, nil
	} else {
		return "", fmt.Errorf("para: lost parameter[%s]", ParaName)
	}
}
