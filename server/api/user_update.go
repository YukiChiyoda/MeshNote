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
	data.Count, err = para.GetInt("count", c)
	if err != nil {
		catch.HandleRequestError(c, err)
		return
	}
	data.Year, _ = strconv.Atoi(time.Now().Format("2006"))
	data.Month, _ = strconv.Atoi(time.Now().Format("01"))
	data.Day, _ = strconv.Atoi(time.Now().Format("02"))
	if err := db.UpdateDiary(data); err != nil {
		catch.HandleServerError(c, err)
		return
	}
	c.String(http.StatusOK, "Successful!")
}
