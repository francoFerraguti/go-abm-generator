package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
)

var router *gin.Engine

func main(){
	router = gin.New()

	router.POST("/create/:project_name", create)

	router.Run(":8000")
}

func create(c *gin.Context) {
	projectName := c.Param("project_name")
	body := frango.ParseRequestBody(c)
	parentFolder := "temp/" + projectName

	config := getConfigStruct(body["config"].(map[string]interface{}))
	models := getModelsArray(body["models"].([]interface{}))

	createFolderStructure(parentFolder)
	createFiles(parentFolder, projectName, config)

    content := gin.H{"success": "true", "description": projectName}
    c.JSON(200, content)
}

func createFolderStructure(parentFolder string) {
	frango.CreateFolder("temp")
	frango.CreateFolder(parentFolder)
	frango.CreateFolder(parentFolder + "/config")
	frango.CreateFolder(parentFolder + "/router")
	frango.CreateFolder(parentFolder + "/controllers")
	frango.CreateFolder(parentFolder + "/models")
	frango.CreateFolder(parentFolder + "/dbhandler")
}

func createFiles(parentFolder string, projectName string, config Config) {
	frango.CreateFile(parentFolder + "/main.go", getFileMainGo(projectName))
	frango.CreateFile(parentFolder + "/config/config.go", getFileConfigGo(projectName, config))
	frango.CreateFile(parentFolder + "/dbhandler/dbhandler.go", getDBHandlerGo(projectName))
}