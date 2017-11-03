package generatorPackageJson

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.PackageJson()

	return template
}
