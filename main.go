package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"

	"github.com/francoFerraguti/go-abm-generator/generator"
	"github.com/francoFerraguti/go-abm-generator/structs"
)

var router *gin.Engine

func main() {
	router = gin.New()

	router.POST("/create", create)

	router.Run(":8000")
}

func create(c *gin.Context) {
	data, err := structs.ParseBody(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	projectName := getProjectName(data.ProjectPath)
	needAuthentication := data.NeedAuthentication
	parentFolder := "temp/" + projectName
	frontendParentFolder := "frontend/" + projectName

	createFolderStructure(parentFolder, needAuthentication)
	createFiles(parentFolder, data.ProjectPath, needAuthentication, data.Config, data.Models)

	createFrontendFolderStructure(frontendParentFolder)
	createFrontendFiles(frontendParentFolder, data.ProjectPath, needAuthentication, data.Config, data.Models)

	c.JSON(200, projectName+" created successfully")
}

func createFrontendFolderStructure(parentFolder string) {
	frango.CreateFolder("frontend")
	frango.CreateFolder(parentFolder)
	frango.CreateFolder(parentFolder + "/public")
	frango.CreateFolder(parentFolder + "/src")
	frango.CreateFolder(parentFolder + "/src/components")
	frango.CreateFolder(parentFolder + "/src/components/App")
	frango.CreateFolder(parentFolder + "/src/components/Layout")
	frango.CreateFolder(parentFolder + "/src/components/Login")
	frango.CreateFolder(parentFolder + "/src/components/MuiTheme")
}

func createFrontendFiles(parentFolder string, projectPath string, needAuthentication bool, config structs.ConfigStruct, models []structs.ModelStruct) {
	frango.CreateFile(parentFolder+"/.gitignore", generator.GetGitIgnore())
	frango.CreateFile(parentFolder+"/package.json", generator.GetPackageJson())
	frango.CreateFile(parentFolder+"/public/index.html", generator.GetIndexHtml())
	frango.CreateFile(parentFolder+"/public/manifest.json", generator.GetManifestJson())

	frango.CreateFile(parentFolder+"/src/registerServiceWorker.js", generator.GetRegisterServiceWorkerJs())
	frango.CreateFile(parentFolder+"/src/logo.svg", generator.GetLogoSvg())
	frango.CreateFile(parentFolder+"/src/index.js", generator.GetIndexJs())
	frango.CreateFile(parentFolder+"/src/index.css", generator.GetIndexCss())
	frango.CreateFile(parentFolder+"/src/App.test.js", generator.GetAppTestJs())

	frango.CreateFile(parentFolder+"/src/components/restricted.js", generator.GetRestrictedJs())
	frango.CreateFile(parentFolder+"/src/components/App/App.css", generator.GetAppCss())
	frango.CreateFile(parentFolder+"/src/components/App/App.js", generator.GetAppJs())
	frango.CreateFile(parentFolder+"/src/components/Layout/Avatar.js", generator.GetAvatarJs())
	frango.CreateFile(parentFolder+"/src/components/Layout/Drawer.js", generator.GetDrawerJs(models))
	frango.CreateFile(parentFolder+"/src/components/Layout/Layout.js", generator.GetLayoutJs(models))
	frango.CreateFile(parentFolder+"/src/components/Layout/Main.js", generator.GetMainJs())
	frango.CreateFile(parentFolder+"/src/components/Login/Login.js", generator.GetLoginJs())
	frango.CreateFile(parentFolder+"/src/components/Login/Api.js", generator.GetLoginApiJs())
	frango.CreateFile(parentFolder+"/src/components/MuiTheme/MuiTheme.js", generator.GetMuiThemeJs())

	for _, model := range models {
		frango.CreateFolder(parentFolder + "/src/components/" + frango.FirstLetterToUpper(model.Name))

		frango.CreateFile(parentFolder+"/src/components/"+frango.FirstLetterToUpper(model.Name)+"/"+frango.FirstLetterToUpper(model.Name)+".js", generator.GetComponentJs(model))
		frango.CreateFile(parentFolder+"/src/components/"+frango.FirstLetterToUpper(model.Name)+"/"+frango.FirstLetterToUpper(model.Name)+"Table.js", generator.GetComponentTableJs(model))
		frango.CreateFile(parentFolder+"/src/components/"+frango.FirstLetterToUpper(model.Name)+"/Endpoints.js", generator.GetComponentEndpointsJs(model))
		frango.CreateFile(parentFolder+"/src/components/"+frango.FirstLetterToUpper(model.Name)+"/Pagination.js", generator.GetComponentPaginationJs())
	}
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

func createFiles(parentFolder string, projectPath string, needAuthentication bool, config structs.ConfigStruct, models []structs.ModelStruct) {
	frango.CreateFile(parentFolder+"/main.go", generator.GetMain(projectPath))
	frango.CreateFile(parentFolder+"/config/config.go", generator.GetConfig(projectPath, needAuthentication, config))
	frango.CreateFile(parentFolder+"/dbhandler/dbhandler.go", generator.GetDBHandler(projectPath, models))
	frango.CreateFile(parentFolder+"/dbhandler/schema.go", generator.GetSchema(models))
	frango.CreateFile(parentFolder+"/structs/structs.go", generator.GetStructs(projectPath, needAuthentication, models))
	frango.CreateFile(parentFolder+"/router/router.go", generator.GetRouter(projectPath, needAuthentication, models))
	frango.CreateFile(parentFolder+"/documentation.md", generator.GetDocumentation(needAuthentication, models))

	createModelsAndControllers(parentFolder, projectPath, needAuthentication, models)

	if needAuthentication {
		frango.CreateFile(parentFolder+"/authentication/authentication.go", generator.GetAuthentication(projectPath, models))
		frango.CreateFile(parentFolder+"/middleware/middleware.go", generator.GetMiddleware(projectPath, models))
		frango.CreateFile(parentFolder+"/controllers/authentication/authentication.go", generator.GetAuthenticationController(projectPath, models))
	}
}

func createModelsAndControllers(parentFolder string, projectPath string, needAuthentication bool, models []structs.ModelStruct) {
	for _, model := range models {
		frango.CreateFolder(parentFolder + "/models/" + frango.FirstLetterToLower(model.Name))
		frango.CreateFolder(parentFolder + "/controllers/" + frango.FirstLetterToLower(model.Name))

		frango.CreateFile(parentFolder+"/models/"+frango.FirstLetterToLower(model.Name)+"/"+frango.FirstLetterToLower(model.Name)+".go", generator.GetModel(projectPath, needAuthentication, model))
		frango.CreateFile(parentFolder+"/controllers/"+frango.FirstLetterToLower(model.Name)+"/"+frango.FirstLetterToLower(model.Name)+".go", generator.GetController(projectPath, model))
	}
}

func getProjectName(projectPath string) string {
	substringArray := strings.Split(projectPath, "/")
	return substringArray[len(substringArray)-1]
}
