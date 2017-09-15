package generatorRouter

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	importArray := []string{"github.com/gin-gonic/gin"}
	endpoints := ""

	ginMiddlewareString, authEndpoints, authImports := getAuthenticationVariables(projectPath, needAuthentication)

	for _, model := range models {
		modelNameLower := frango.FirstLetterToLower(model.Name)
		modelNameUpper := frango.FirstLetterToUpper(model.Name)

		importArray = append(importArray, projectPath+"/controllers/"+modelNameLower)

		endpoints += getEndpoints(model.Fields, modelNameLower, modelNameUpper)
	}

	importArray = append(importArray, authImports)

	template := templates.RouterGo()

	imports := common.GetImportsWithArray(importArray)

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&AUTHENTICATION_ENDPOINTS&&", authEndpoints, -1)
	fileContent = strings.Replace(fileContent, "&&GIN_MIDDLEWARE_STRING&&", ginMiddlewareString, -1)
	fileContent = strings.Replace(fileContent, "&&ENDPOINTS&&", endpoints, -1)

	return fileContent
}

func getEndpoints(fields []structs.FieldStruct, modelNameLower, modelNameUpper string) string {
	endpoints := ""

	endpoints += "		api.POST(`/" + modelNameLower + "`, " + modelNameLower + ".Create)\n"

	for _, field := range fields {
		if !field.Unique {
			continue
		}

		endpoints += "		api.GET(`/" + modelNameLower + "/" + modelNameLower + "/:" + modelNameLower + "`, " + modelNameLower + ".GetBy" + modelNameUpper + ")\n"
		endpoints += "		api.PUT(`/" + modelNameLower + "/" + modelNameLower + "/:" + modelNameLower + "`, " + modelNameLower + ".UpdateBy" + modelNameUpper + ")\n"
		endpoints += "		api.DELETE(`/" + modelNameLower + "/" + modelNameLower + "/:" + modelNameLower + "`, " + modelNameLower + ".DeleteBy" + modelNameUpper + ")\n"
	}

	return endpoints
}

func getAuthenticationVariables(projectPath string, needAuthentication bool) (string, string, string) {
	ginMiddlewareString := ""
	authEndpoints := ""
	authImports := ""

	if needAuthentication {
		ginMiddlewareString = ", middleware.ValidateToken()"

		authEndpoints = `public := router.Group("/")
	{
		public.POST("/signup", authentication.Signup)
		public.POST("/login", authentication.Login)
	}`

		authImports = projectPath + `/middleware"
	"` + projectPath + `/controllers/authentication`
	}

	return ginMiddlewareString, authEndpoints, authImports
}
