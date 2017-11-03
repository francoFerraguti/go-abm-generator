package generatorManifestJson

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.ManifestJson()

	return template
}
