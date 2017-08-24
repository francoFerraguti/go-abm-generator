package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
)

var router *gin.Engine

func main() {
	router = gin.New()

	router.POST("/create/:project_name", create)

	router.Run(":8000")
}

func create(c *gin.Context) {
	projectName := c.Param("project_name")
	parentFolder := "temp/" + projectName

	data, err := parseBody(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	createFolderStructure(parentFolder)
	createFiles(parentFolder, projectName, data.Config, data.Models)

	c.JSON(200, projectName+" created successfully")
}

func createFolderStructure(parentFolder string) {
	frango.CreateFolder("temp")
	frango.CreateFolder(parentFolder)
	frango.CreateFolder(parentFolder + "/config")
	frango.CreateFolder(parentFolder + "/router")
	frango.CreateFolder(parentFolder + "/controllers")
	frango.CreateFolder(parentFolder + "/models")
	frango.CreateFolder(parentFolder + "/dbhandler")
	frango.CreateFolder(parentFolder + "/structs")
}

func createFiles(parentFolder string, projectName string, config ConfigStruct, models []ModelStruct) {
	frango.CreateFile(parentFolder+"/main.go", getFileMainGo(projectName))
	frango.CreateFile(parentFolder+"/config/config.go", getFileConfigGo(projectName, config))
	frango.CreateFile(parentFolder+"/dbhandler/dbhandler.go", getFileDBHandlerGo(projectName))
	frango.CreateFile(parentFolder+"/structs/structs.go", getFileStructsGo(projectName, models))

	createModelsAndControllers(parentFolder, models)
}

func createModelsAndControllers(parentFolder string, models []ModelStruct) {
	for _, model := range models {
		frango.CreateFolder(parentFolder + "/models/" + strings.ToLower(model.Name))
		frango.CreateFolder(parentFolder + "/controllers/" + strings.ToLower(model.Name))

		frango.CreateFile(parentFolder+"/models/"+strings.ToLower(model.Name)+"/"+strings.ToLower(model.Name)+".go", getFileModelGo(parentFolder, model))
	}
}
