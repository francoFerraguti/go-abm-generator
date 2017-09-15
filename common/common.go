package common

import (
	"strings"
)

func GetImports(files ...string) string {
	s := `
import (
	&&IMPORTS&&
)`
	importsString := getImportsString(files)

	s = strings.Replace(s, "&&IMPORTS&&", importsString, -1)

	return strings.Trim(s, "\n")
}

func GetImportsWithArray(files []string) string {
	s := `
import (
	&&IMPORTS&&
)`
	importsString := getImportsString(files)

	s = strings.Replace(s, "&&IMPORTS&&", importsString, -1)

	return strings.Trim(s, "\n")
}

func getImportsString(files []string) string {
	for key, _ := range files {
		files[key] = "\t" + `"` + files[key] + `"`
	}

	importsString := ""

	for _, val := range files {
		importsString += val + "\n"
	}

	importsString = strings.TrimLeft(importsString, "\t")
	importsString = strings.Trim(importsString, "\n")

	return importsString
}
