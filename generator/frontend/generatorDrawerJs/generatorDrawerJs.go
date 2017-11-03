package generatorDrawerJs

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"strconv"
	"strings"
)

func Get(models []structs.ModelStruct) string {
	template := templates.DrawerJs()

	routesList := getRoutesList(models)

	fileContent := strings.Replace(template, "&&ROUTES_LIST&&", routesList, -1)

	return fileContent
}

func getRoutesList(models []structs.ModelStruct) string {
	routesList := ""

	for key, model := range models {
		routesList += "{this.createItem(" + strconv.Itoa(key) + ", '" + model.Name + "', '" + model.Name + "', '/home/" + model.Name + "', <IconAssignmentTurnedIn/>, '')}\n"
	}

	return routesList
}
