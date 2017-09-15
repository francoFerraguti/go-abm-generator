package common

import(
	"strings"
)

func GetImports(projectPath string, files... string) string {
	s := `
import (
	&&IMPORTS&&
)`
	importsString := getImportsString(projectPath, files)

	s = strings.Replace(s, "&&IMPORTS&&", importsString, -1)

	return strings.Trim(s, "\n")
}

func getImportsString(projectPath string, files []string) string {
	for key, _ := range files {
		files[key] = "\t" + `"` + projectPath + "/" + files[key] + `"`
	}

	importsString := ""

	for _, val := range files {
		importsString += val + "\n"
	}
	
	importsString = strings.TrimLeft(importsString, "\t")
	importsString = strings.Trim(importsString, "\n")	

	return importsString
}