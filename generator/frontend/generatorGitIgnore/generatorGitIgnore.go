package generatorGitIgnore

import (
	"github.com/francoFerraguti/go-abm-generator/templates"
)

func Get() string {
	template := templates.GitIgnore()

	return template
}
