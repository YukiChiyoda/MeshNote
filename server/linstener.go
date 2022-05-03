package main

import (
	"MeshNote/server/api"

	"github.com/gin-gonic/gin"
)

func ListenerBuild() {
	router := gin.Default()
	router.GET("/", api.Root)
	router.GET("/query.go", api.Query)
	router.POST("/read.go", api.Read)
	router.POST("/create.go", api.Create)
	router.POST("/write.go", api.Write)
	router.POST("/move.go", api.Move)
	router.POST("/drop.go", api.Drop)
	router.Run(":3000")
}
