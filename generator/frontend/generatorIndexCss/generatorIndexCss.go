package generatorIndexCss

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.IndexCss()

	return template
}
