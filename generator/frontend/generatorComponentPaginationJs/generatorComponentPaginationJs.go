package generatorComponentPaginationJs

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.ComponentPaginationJs()

	return template
}
