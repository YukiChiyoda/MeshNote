package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	targetUser := db.DEV_USER_ID
	/*
		[DEV USER]
		targetUser, err := para.GetInt("user", c)
		if err != nil {
			catch.HandleRequestError(c, err)
			return
		}
		db.IfUserExist()
		...
	*/
	targetCount, err := para.GetInt("count", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetYear, _ := strconv.Atoi(time.Now().Format("2006"))
	targetMonth, _ := strconv.Atoi(time.Now().Format("01"))
	targetDay, _ := strconv.Atoi(time.Now().Format("02"))
	if err := db.UpdateDiary(targetUser, targetYear, targetMonth, targetDay, targetCount); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, "Successful!")
}
