package generator

import (
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorAuthentication"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorAuthenticationController"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorConfig"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorController"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorDBHandler"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorDocumentation"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorMain"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorMiddleware"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorModel"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorRouter"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorSchema"
	"github.com/francoFerraguti/go-abm-generator/generator/backend/generatorStructs"

	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorAppCss"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorAppJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorAppTestJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorAvatarJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorComponentEndpointsJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorComponentJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorComponentPaginationJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorComponentTableJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorDrawerJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorGitIgnore"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorIndexCss"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorIndexHtml"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorIndexJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorLayoutJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorLoginApiJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorLoginJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorLogoSvg"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorMainJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorManifestJson"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorMuiThemeJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorPackageJson"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorRegisterServiceWorkerJs"
	"github.com/francoFerraguti/go-abm-generator/generator/frontend/generatorRestrictedJs"

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

func GetStructs(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	return generatorStructs.Get(projectPath, needAuthentication, models)
}

func GetDBHandler(projectPath string, models []structs.ModelStruct) string {
	return generatorDBHandler.Get(projectPath, models)
}

func GetSchema(models []structs.ModelStruct) string {
	return generatorSchema.Get(models)
}

func GetDocumentation(needAuthentication bool, models []structs.ModelStruct) string {
	return generatorDocumentation.Get(needAuthentication, models)
}

func GetAuthentication(projectPath string, models []structs.ModelStruct) string {
	return generatorAuthentication.Get(projectPath, models)
}

func GetAuthenticationController(projectPath string, models []structs.ModelStruct) string {
	return generatorAuthenticationController.Get(projectPath, models)
}

func GetModel(projectPath string, needAuthentication bool, model structs.ModelStruct) string {
	return generatorModel.Get(projectPath, needAuthentication, model)
}

func GetController(projectPath string, model structs.ModelStruct) string {
	return generatorController.Get(projectPath, model)
}

//

func GetAppCss() string {
	return generatorAppCss.Get()
}

func GetAppJs() string {
	return generatorAppJs.Get()
}
func GetAppTestJs() string {
	return generatorAppTestJs.Get()
}
func GetAvatarJs() string {
	return generatorAvatarJs.Get()
}
func GetComponentEndpointsJs(model structs.ModelStruct) string {
	return generatorComponentEndpointsJs.Get(model)
}
func GetComponentJs(model structs.ModelStruct) string {
	return generatorComponentJs.Get(model)
}
func GetComponentPaginationJs() string {
	return generatorComponentPaginationJs.Get()
}
func GetComponentTableJs(model structs.ModelStruct) string {
	return generatorComponentTableJs.Get(model)
}
func GetDrawerJs(models []structs.ModelStruct) string {
	return generatorDrawerJs.Get(models)
}
func GetGitIgnore() string {
	return generatorGitIgnore.Get()
}
func GetIndexCss() string {
	return generatorIndexCss.Get()
}
func GetIndexHtml() string {
	return generatorIndexHtml.Get()
}
func GetIndexJs() string {
	return generatorIndexJs.Get()
}
func GetLayoutJs(models []structs.ModelStruct) string {
	return generatorLayoutJs.Get(models)
}
func GetLoginApiJs() string {
	return generatorLoginApiJs.Get()
}
func GetLoginJs() string {
	return generatorLoginJs.Get()
}
func GetLogoSvg() string {
	return generatorLogoSvg.Get()
}
func GetMainJs() string {
	return generatorMainJs.Get()
}
func GetManifestJson() string {
	return generatorManifestJson.Get()
}
func GetMuiThemeJs() string {
	return generatorMuiThemeJs.Get()
}
func GetPackageJson() string {
	return generatorPackageJson.Get()
}
func GetRegisterServiceWorkerJs() string {
	return generatorRegisterServiceWorkerJs.Get()
}
func GetRestrictedJs() string {
	return generatorRestrictedJs.Get()
}
