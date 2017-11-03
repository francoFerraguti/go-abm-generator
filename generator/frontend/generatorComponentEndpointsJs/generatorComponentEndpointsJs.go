package generatorComponentEndpointsJs

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(model structs.ModelStruct) string {
	template := templates.ComponentEndpointsJs()

	createFunction := getCreateFunction(model)
	editFunction := getEditFunction(model)
	deleteFunction := getDeleteFunction(model)
	getAllFunction := getGetAllFunction(model)

	fileContent := strings.Replace(template, "&&COMPONENT_NAME&&", model.Name, -1)
	fileContent = strings.Replace(fileContent, "&&CREATE_FUNCTION&&", createFunction, -1)
	fileContent = strings.Replace(fileContent, "&&EDIT_FUNCTION&&", editFunction, -1)
	fileContent = strings.Replace(fileContent, "&&DELETE_FUNCTION&&", deleteFunction, -1)
	fileContent = strings.Replace(fileContent, "&&GET_ALL_FUNCTION&&", getAllFunction, -1)

	return fileContent
}

func getCreateFunction(model structs.ModelStruct) string {
	output := ""

	output += "create" + model.Name + "(" + frango.FirstLetterToLower(model.Name) + ") {\n"
	output += "axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');\n"
	output += "return axios.post(URL_BASE + '/" + frango.FirstLetterToLower(model.Name) + "', " + frango.FirstLetterToLower(model.Name) + ", {});\n"
	output += "},\n"

	return output
}

func getEditFunction(model structs.ModelStruct) string {
	output := ""

	output += "edit" + model.Name + "(id" + model.Name + ", " + frango.FirstLetterToLower(model.Name) + ") {\n"
	output += "axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');\n"
	output += "return axios.put(URL_BASE + '/" + frango.FirstLetterToLower(model.Name) + "/id/' + id" + model.Name + ", " + frango.FirstLetterToLower(model.Name) + ", {});\n"
	output += "},\n"

	return output
}

func getDeleteFunction(model structs.ModelStruct) string {
	output := ""

	output += "delete" + model.Name + "(id" + model.Name + ") {\n"
	output += "axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');\n"
	output += "return axios.delete(URL_BASE + '/" + frango.FirstLetterToLower(model.Name) + "/id/' + id" + model.Name + ", {});\n"
	output += "},\n"

	return output
}

func getGetAllFunction(model structs.ModelStruct) string {
	output := ""

	output += "getAll" + model.Name + "() {\n"
	output += "axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');\n"
	output += "return axios.get(URL_BASE + '/" + frango.FirstLetterToLower(model.Name) + "/list', {});\n"
	output += "},\n"

	return output
}
