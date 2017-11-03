package generatorAppCss

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.AppCss()

	return template
}
