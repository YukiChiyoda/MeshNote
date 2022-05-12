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
	var data db.Diary
	var err error
	data.User = db.DEV_USER_ID
	/*
		[DEV USER]
		data.User, err := para.GetInt("user", c)
		if err != nil {
			catch.HandleRequestError(c, err)
			return
		}
		db.IfUserExist()
		...
	*/
	data.Year, err = para.GetInt("year", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	data.Month, err = para.GetInt("month", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	data.Day, err = para.GetInt("day", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	temp, err := db.QueryDiary(data)
	if err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, strconv.Itoa(temp))
}
