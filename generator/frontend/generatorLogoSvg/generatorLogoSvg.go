package generatorLogoSvg

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.LogoSvg()

	return template
}
