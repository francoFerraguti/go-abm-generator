package generatorIndexHtml

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.IndexHtml()

	return template
}
