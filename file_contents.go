package main

import (
	"github.com/liteByte/frango"
	"strings"
)

func getFileMainGo(projectName string) string {
	return `package main

import (
	"` + projectName + `/dbhandler"
	"` + projectName + `/router"
)

func main() {
	dbhandler.ConnectToDatabase()
	router.ConfigureRouter()
	router.CreateRouter()
	router.RunRouter()
}
	`
}

func getFileConfigGo(projectName string, config ConfigStruct) string {
	return `package config

type Config struct {
	DB_TYPE		string
	DB_USERNAME	string
	DB_PASSWORD	string
	DB_HOST		string
	DB_PORT		string
	DB_NAME 	string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		config := newConfigLocal()
		instance = &config
	}
	return instance
}

func newConfigLocal() Config {
	return Config{
		DB_TYPE:       	"` + config.DB_TYPE + `",
		DB_USERNAME:    "` + config.DB_USERNAME + `",
		DB_PASSWORD:    "` + config.DB_PASSWORD + `",
		DB_HOST:      	"` + config.DB_HOST + `",
		DB_PORT:       	"` + config.DB_PORT + `",
		DB_NAME:       	"` + config.DB_NAME + `",
	}
}
	`
}

func getFileDBHandlerGo(projectName string) string {
	return `package dbhandler

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/liteByte/frango"
)

var db *sql.DB

func ConnectToDatabase() {
    var err error
	
	db, err = sql.Open(config.GetConfig().DB_TYPE, config.GetConfig().DB_USERNAME + ":" + config.GetConfig().DB_PASSWORD + "@tcp(" + config.GetConfig().DB_HOST + ":" + config.GetConfig().DB_PORT + ")/" + config.GetConfig().DB_NAME)
	frango.PrintErr(err)
    
    err = db.Ping()
    frango.PrintErr(err)
}

func GetDatabase() *sql.DB {
    return db
}
	`
}

func getFileStructsGo(projectName string, models []ModelStruct) string {
	structsString := ""
	for _, model := range models {
		structsString += "type " + model.Name + "Struct struct {\n"
		for _, field := range model.Fields {
			structsString += "	" + frango.FirstLetterToUpper(field.Name) + " " + field.Type + "\n"
		}
		structsString += "}\n\n"
	}

	return `package structs

` + structsString
}

func getFileModelGo(projectName string, model ModelStruct) string {
	getByString := ""
	deleteByString := ""

	for _, field := range model.Fields {
		if !field.Unique {
			continue
		}

		getByString += modelGetBy(model, field)
		deleteByString += modelDeleteBy(model, field)
	}

	return `package ` + strings.ToLower(model.Name) + `

import (
	"` + projectName + `/dbhandler"
    "` + projectName + `/structs"
)

func Update(` + frango.FirstLetterToLower(model.Name) + " " + model.Name + `Struct` + `) error {

}

` + modelCreate(model) + getByString + deleteByString
}

// Internal functions
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
		fieldsStructList += frango.FirstLetterToLower(model.Name) + "." + frango.FirstLetterToLower(field.Name) + ", "
	}
	fieldsList = frango.RemoveLastCharacters(fieldsList, 2)
	questionMarksList = frango.RemoveLastCharacters(questionMarksList, 2)
	fieldsStructList = frango.RemoveLastCharacters(fieldsStructList, 2)

	createString += "func Create(" + frango.FirstLetterToLower(model.Name) + " " + model.Name + "Struct" + ") error {\n"
	createString += "	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO " + model.Name + " (" + fieldsList + ") VALUES(" + questionMarksList + ")`, " + fieldsStructList + ")\n"
	createString += "	return err\n"
	createString += "}\n\n"

	return createString
}

func modelGetBy(model ModelStruct, field FieldStruct) string {
	getByString := ""
	fieldsList := ""
	fieldsListAmpersand := ""

	getByString += "func GetBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + " " + field.Type + ") (" + model.Name + "Struct, error) {\n"
	getByString += "	var " + frango.FirstLetterToLower(model.Name) + " " + model.Name + "Struct\n\n"
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
	getByString += "		return " + model.Name + "Struct{}, err\n"
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
