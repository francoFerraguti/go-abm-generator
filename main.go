package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
)

var router *gin.Engine

func main() {
	router = gin.New()

	router.POST("/create", create)

	router.Run(":8000")
}

func create(c *gin.Context) {
	data, err := parseBody(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	projectName := getProjectName(data.ProjectPath)
	needAuthentication := data.NeedAuthentication
	parentFolder := "temp/" + projectName

	createFolderStructure(parentFolder, needAuthentication)
	createFiles(parentFolder, data.ProjectPath, needAuthentication, data.Config, data.Models)

	c.JSON(200, projectName+" created successfully")
}

func createFolderStructure(parentFolder string, needAuthentication bool) {
	frango.CreateFolder("temp")
	frango.CreateFolder(parentFolder)
	frango.CreateFolder(parentFolder + "/config")
	frango.CreateFolder(parentFolder + "/router")
	frango.CreateFolder(parentFolder + "/controllers")
	frango.CreateFolder(parentFolder + "/models")
	frango.CreateFolder(parentFolder + "/dbhandler")
	frango.CreateFolder(parentFolder + "/structs")

	if needAuthentication {
		frango.CreateFolder(parentFolder + "/authentication")
		frango.CreateFolder(parentFolder + "/middleware")
		frango.CreateFolder(parentFolder + "/controllers/authentication")
	}
}

func createFiles(parentFolder string, projectPath string, needAuthentication bool, config ConfigStruct, models []ModelStruct) {
	frango.CreateFile(parentFolder+"/main.go", getFileMainGo(projectPath))
	frango.CreateFile(parentFolder+"/config/config.go", getFileConfigGo(projectPath, needAuthentication, config))
	frango.CreateFile(parentFolder+"/dbhandler/dbhandler.go", getFileDBHandlerGo(projectPath, models))
	frango.CreateFile(parentFolder+"/dbhandler/schema.go", getFileDBSchemaGo(models))
	frango.CreateFile(parentFolder+"/structs/structs.go", getFileStructsGo(projectPath, needAuthentication, models))
	frango.CreateFile(parentFolder+"/router/router.go", getFileRouterGo(projectPath, needAuthentication, models))

	createModelsAndControllers(parentFolder, projectPath, needAuthentication, models)

	if needAuthentication {
		frango.CreateFile(parentFolder+"/authentication/authentication.go", getFileAuthenticationGo(projectPath, models))
		frango.CreateFile(parentFolder+"/middleware/middleware.go", getFileMiddlewareGo(projectPath, models))
		frango.CreateFile(parentFolder+"/controllers/authentication/authentication.go", getFileControllerAuthenticationGo(projectPath, models))
	}
}

func createModelsAndControllers(parentFolder string, projectPath string, needAuthentication bool, models []ModelStruct) {
	for _, model := range models {
		frango.CreateFolder(parentFolder + "/models/" + strings.ToLower(model.Name))
		frango.CreateFolder(parentFolder + "/controllers/" + strings.ToLower(model.Name))

		frango.CreateFile(parentFolder+"/models/"+strings.ToLower(model.Name)+"/"+strings.ToLower(model.Name)+".go", getFileModelGo(projectPath, needAuthentication, model))
		frango.CreateFile(parentFolder+"/controllers/"+strings.ToLower(model.Name)+"/"+strings.ToLower(model.Name)+".go", getFileControllerGo(projectPath, model))
	}
}

func getProjectName(projectPath string) string {
	substringArray := strings.Split(projectPath, "/")
	return substringArray[len(substringArray)-1]
}
