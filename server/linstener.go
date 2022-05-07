package main

import (
	"MeshNote/server/api"

	"github.com/gin-gonic/gin"
)

func ListenerBuild() {
	router := gin.Default()
	router.POST("/", api.Root)
	router.POST("/query.go", api.Query)
	router.POST("/read.go", api.Read)
	router.POST("/write.go", api.Write)
	router.POST("/create.go", api.Create)
	router.POST("/delete.go", api.Delete)
	router.POST("/move.go", api.Move)
	router.Run(":3000")
}
