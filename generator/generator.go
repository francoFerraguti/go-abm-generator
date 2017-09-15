package generator

import (
	"github.com/francoFerraguti/go-abm-generator/generator/generatorConfig"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorMain"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorMiddleware"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorRouter"
	"github.com/francoFerraguti/go-abm-generator/structs"
)

func GetMain(projectPath string) string {
	return generatorMain.Get(projectPath)
}

func GetMiddleware(projectPath string, models []structs.ModelStruct) string {
	return generatorMiddleware.Get(projectPath, models)
}

func GetConfig(projectPath string, needAuthentication bool, config structs.ConfigStruct) string {
	return generatorConfig.Get(projectPath, needAuthentication, config)
}

func GetRouter(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	return generatorRouter.Get(projectPath, needAuthentication, models)
}
