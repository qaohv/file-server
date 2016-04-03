package main

import "github.com/gin-gonic/gin"

func makeRouter(pathToFolder string) *gin.Engine {
	api := makeNewApi(pathToFolder)

	engine := gin.Default()
	v1 := engine.Group("/v1")
	{
		v1.GET("/files", api.ListFilesHandler)
		v1.Static("/file", pathToFolder)
	}

	return engine
}
