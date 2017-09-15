package generatorMain

import (
	"strings"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/francoFerraguti/go-abm-generator/common"
)

func Get(projectPath string) string {
	template := templates.MainGo()

	imports := common.GetImports(projectPath, "dbhandler", "router")

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)

	return fileContent
}

