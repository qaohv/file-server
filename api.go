package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Api struct {
	PathToFolder string
}

func makeNewApi(pathToFolder string) *Api {
	return &Api{PathToFolder: pathToFolder}
}

func (api *Api) ListFilesHandler(c *gin.Context) {
	files, err := ioutil.ReadDir(api.PathToFolder)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"path":  api.PathToFolder,
		}).Error("ReadDir error")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	names := make([]string, len(files))

	for index, file := range files {
		names[index] = file.Name()
	}

	c.JSON(http.StatusOK, gin.H{"files": names})
}
