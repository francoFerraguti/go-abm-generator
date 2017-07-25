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

	createFolderStructure(parentFolder)
	createFiles(parentFolder, projectName, config)

    content := gin.H{"success": "true", "description": projectName}
    c.JSON(200, content)
}

func getConfigStruct(config map[string]interface{}) Config {
	return Config {
		DB_TYPE: config["db_type"].(string),
		DB_USERNAME: config["db_username"].(string),
		DB_PASSWORD: config["db_password"].(string),
		DB_HOST: config["db_host"].(string),
		DB_PORT: config["db_port"].(string),
		DB_NAME: config["db_name"].(string),
	}
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
}