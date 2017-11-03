package generatorIndexJs

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.IndexJs()

	return template
}
