package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	PATH_TO_TEST_FOLDER = "./testfolder"
)

type FileListResult struct {
	Files []string `json:"files"`
}

func makeRequest(router *gin.Engine, method, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func getTestFolderFilenames(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"path":  path,
		}).Fatal("Read dir error")
	}
	names := make([]string, len(files))

	for index, file := range files {
		names[index] = file.Name()
	}
	return names
}

func TestListFiles(t *testing.T) {
	r := makeRouter(PATH_TO_TEST_FOLDER)

	response := makeRequest(r, "GET", "/v1/files")

	fileListResult := new(FileListResult)
	json.NewDecoder(response.Body).Decode(fileListResult)

	names := getTestFolderFilenames(PATH_TO_TEST_FOLDER)

	assert.Equal(t, http.StatusOK, response.Code, "Code should be OK")
	assert.Equal(t, fileListResult.Files, names, "Filenames should be equal")
}
