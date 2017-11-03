package generatorLayoutJs

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"strings"
)

func Get(models []structs.ModelStruct) string {
	template := templates.LayoutJs()

	importList := getImportList(models)
	routeList := getRouteList(models)

	fileContent := strings.Replace(template, "&&IMPORT_LIST&&", importList, -1)
	fileContent = strings.Replace(fileContent, "&&ROUTE_LIST&&", routeList, -1)

	return fileContent
}

func getImportList(models []structs.ModelStruct) string {
	importList := ""

	for _, model := range models {
		importList += "import " + model.Name + " from '../" + model.Name + "/" + model.Name + ".js'\n"
	}

	return importList
}

func getRouteList(models []structs.ModelStruct) string {
	routeList := ""

	for _, model := range models {
		routeList += "<Route path='/home/" + model.Name + "' render={() => <" + model.Name + "/>}/>\n"
	}

	return routeList
}
