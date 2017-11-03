package generatorLoginJs

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.LoginJs()

	return template
}
