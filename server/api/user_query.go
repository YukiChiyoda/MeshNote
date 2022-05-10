package api

import (
	"MeshNote/server/api/catch"
	"MeshNote/server/api/para"
	"MeshNote/server/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserQuery(c *gin.Context) {
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
	targetYear, err := para.GetInt("year", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetMonth, err := para.GetInt("month", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	targetDay, err := para.GetInt("day", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	temp, err := db.QueryDiary(targetUser, targetYear, targetMonth, targetDay)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, strconv.Itoa(temp))
}
