package main

import (
	"github.com/liteByte/frango"
	"strings"
)

func getFileMainGo(projectName string) string {
	return `package main

import (
	` + projectName + `/dbhandler
	` + projectName + `/router
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
		DB_NAME:       	"` + config.DB_NAME + `"
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
	return `package ` + strings.ToLower(projectName) + `

import (
	` + projectName + `/dbhandler
	` + projectName + `/structs
)

func Create()
	`
}
