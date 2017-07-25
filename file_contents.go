package main

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

func getFileConfigGo(projectName string, config Config) string {
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
		DB_TYPE:       	"`+ config.DB_TYPE +`",
		DB_USERNAME:    "`+ config.DB_USERNAME +`",
		DB_PASSWORD:    "`+ config.DB_PASSWORD +`",
		DB_HOST:      	"`+ config.DB_HOST +`",
		DB_PORT:       	"`+ config.DB_PORT +`",
		DB_NAME:       	"`+ config.DB_NAME +`"
	}
}
	`
}