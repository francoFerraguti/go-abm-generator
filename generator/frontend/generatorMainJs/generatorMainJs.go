package generatorMainJs

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.MainJs()

	return template
}
