package generator

import (
	"github.com/francoFerraguti/go-abm-generator/generator/generatorMain"
)

func GetMain(projectPath string) string {
	return generatorMain.Get(projectPath)
}