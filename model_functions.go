package main

import (
	"github.com/liteByte/frango"
)

func modelCreate(model ModelStruct) string {
	createString := ""
	fieldsList := ""
	questionMarksList := ""
	fieldsStructList := ""

	for _, field := range model.Fields {
		if field.AutoGenerated {
			continue
		}

		fieldsList += frango.FirstLetterToLower(field.Name) + ", "
		questionMarksList += "?, "
		fieldsStructList += frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToUpper(field.Name) + ", "
	}
	fieldsList = frango.RemoveLastCharacters(fieldsList, 2)
	questionMarksList = frango.RemoveLastCharacters(questionMarksList, 2)
	fieldsStructList = frango.RemoveLastCharacters(fieldsStructList, 2)

	createString += "func Create(" + frango.FirstLetterToLower(model.Name) + " structs." + model.Name + "Struct" + ") error {\n"
	createString += "	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO " + model.Name + " (" + fieldsList + ") VALUES(" + questionMarksList + ")`, " + fieldsStructList + ")\n"
	createString += "	return err\n"
	createString += "}\n\n"

	return createString
}

func modelUpdateBy(model ModelStruct, field FieldStruct) string {
	updateString := ""
	fieldsList := ""
	parametersList := ""

	for _, secondaryField := range model.Fields {
		if field == secondaryField || secondaryField.AutoGenerated {
			continue
		}

		fieldsList += frango.FirstLetterToLower(secondaryField.Name) + " = ?, "
		parametersList += frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToUpper(secondaryField.Name) + ", "
	}
	parametersList += frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToUpper(field.Name)

	fieldsList = frango.RemoveLastCharacters(fieldsList, 2)

	updateString += "func UpdateBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(model.Name) + " structs." + model.Name + "Struct" + ") error {\n"
	updateString += "	_, err := dbhandler.GetDatabase().Exec(`UPDATE " + model.Name + " SET " + fieldsList + " WHERE " + frango.FirstLetterToLower(field.Name) + " = ?" + "`, " + parametersList + ")\n"
	updateString += "	return err\n"
	updateString += "}\n\n"

	return updateString
}

func modelGetBy(model ModelStruct, field FieldStruct) string {
	getByString := ""
	fieldsList := ""
	fieldsListAmpersand := ""

	getByString += "func GetBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + " " + field.Type + ") (structs." + model.Name + "Struct, error) {\n"
	getByString += "	var " + frango.FirstLetterToLower(model.Name) + " structs." + model.Name + "Struct\n"
	getByString += "	" + frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToUpper(field.Name) + " = " + frango.FirstLetterToLower(field.Name) + "\n\n"
	for _, secondaryField := range model.Fields {
		if field == secondaryField {
			continue
		}
		fieldsList += frango.FirstLetterToLower(secondaryField.Name) + ", "
		fieldsListAmpersand += "&" + frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToUpper(secondaryField.Name) + ", "
	}
	fieldsList = frango.RemoveLastCharacters(fieldsList, 2)
	fieldsListAmpersand = frango.RemoveLastCharacters(fieldsListAmpersand, 2)

	getByString += "	err := dbhandler.GetDatabase().QueryRow(`SELECT " + fieldsList + " FROM " + model.Name + " WHERE " + frango.FirstLetterToLower(field.Name) + " = ?`, " + frango.FirstLetterToLower(field.Name) + ").Scan(" + fieldsListAmpersand + ")\n"
	getByString += "	if err != nil {\n"
	getByString += "		return structs." + model.Name + "Struct{}, err\n"
	getByString += "	}\n\n"
	getByString += "	return " + frango.FirstLetterToLower(model.Name) + ", nil\n"
	getByString += "}\n\n"

	return getByString
}

func modelDeleteBy(model ModelStruct, field FieldStruct) string {
	deleteByString := ""

	deleteByString += "func DeleteBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + " " + field.Type + ") error {\n"
	deleteByString += "    _, err := dbhandler.GetDatabase().Query(`DELETE FROM " + model.Name + " WHERE " + frango.FirstLetterToLower(field.Name) + " = ?`, " + frango.FirstLetterToLower(field.Name) + ")\n"
	deleteByString += "    return err\n"
	deleteByString += "}\n\n"

	return deleteByString
}
