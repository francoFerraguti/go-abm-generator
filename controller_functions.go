package main

import (
	"github.com/liteByte/frango"
	"strings"
)

func controllerCreate(model ModelStruct) string {
	createString := ""

	createString += "func Create(c *gin.Context) {\n"
	createString += "	" + frango.FirstLetterToLower(model.Name) + "Struct, err := structs.ParseBodyInto" + model.Name + "Struct(c.Request.Body)\n"
	createString += "	if err != nil {\n"
	createString += "		c.JSON(400, err.Error())\n"
	createString += "		return\n"
	createString += "	}\n\n"
	createString += "	if err = " + strings.ToLower(model.Name) + ".Create(" + frango.FirstLetterToLower(model.Name) + "Struct); err != nil {\n"
	createString += "		c.JSON(500, err.Error())\n"
	createString += "		return\n"
	createString += "	}\n\n"
	createString += "	c.JSON(200, `" + model.Name + " created successfully`)\n"
	createString += "}\n\n"

	return createString
}

func controllerUpdateBy(model ModelStruct, field FieldStruct) string {
	updateString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	updateString += "func UpdateBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	updateString += "	" + frango.FirstLetterToLower(model.Name) + "Struct, err := structs.ParseBodyInto" + model.Name + "Struct(c.Request.Body)\n"
	updateString += "	if err != nil {\n"
	updateString += "		c.JSON(400, err.Error())\n"
	updateString += "		return\n"
	updateString += "	}\n\n"
	updateString += "	" + frango.FirstLetterToLower(model.Name) + "Struct." + frango.FirstLetterToUpper(field.Name) + " = " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	updateString += "	if err = " + strings.ToLower(model.Name) + ".UpdateBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(model.Name) + "Struct); err != nil {\n"
	updateString += "		c.JSON(500, err.Error())\n"
	updateString += "		return\n"
	updateString += "	}\n\n"
	updateString += "	c.JSON(200, `" + model.Name + " updated successfully`)\n"
	updateString += "}\n\n"

	return updateString
}

func controllerGetBy(model ModelStruct, field FieldStruct) string {
	getByString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	getByString += "func GetBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	getByString += "    " + frango.FirstLetterToLower(field.Name) + " := " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	getByString += "    " + frango.FirstLetterToLower(model.Name) + "Struct, err := " + strings.ToLower(model.Name) + ".GetBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + ")\n"
	getByString += "    if err != nil {\n"
	getByString += "        c.JSON(500, err.Error())\n"
	getByString += "        return\n"
	getByString += "    }\n\n"
	getByString += "    " + frango.FirstLetterToLower(model.Name) + "JSON, _ := json.Marshal(" + frango.FirstLetterToLower(model.Name) + "Struct)\n"
	getByString += "    c.JSON(200, json.RawMessage(string(" + frango.FirstLetterToLower(model.Name) + "JSON)))\n"
	getByString += "}\n\n"

	return getByString
}

func controllerDeleteBy(model ModelStruct, field FieldStruct) string {
	deleteByString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	deleteByString += "func DeleteBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	deleteByString += "    " + frango.FirstLetterToLower(field.Name) + " := " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	deleteByString += "    if err := " + strings.ToLower(model.Name) + ".DeleteBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + "); err != nil {\n"
	deleteByString += "        c.JSON(500, err.Error())\n"
	deleteByString += "        return\n"
	deleteByString += "    }\n\n"
	deleteByString += "    c.JSON(200, `" + model.Name + " deleted successfully`)\n"
	deleteByString += "}\n\n"

	return deleteByString
}
