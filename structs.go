package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type DataStruct struct {
	Config ConfigStruct
	Models []ModelStruct
}

type ConfigStruct struct {
	DB_TYPE     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

type ModelStruct struct {
	Name   string
	Fields []FieldStruct
}

type FieldStruct struct {
	Name     string
	Type     string
	Required bool
	Default  string
	Unique   bool
}

func parseBody(body io.ReadCloser) (DataStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)

	dataStruct := DataStruct{}
	err := json.Unmarshal(bodyBytes, &dataStruct)

	return dataStruct, err
}

/*
{
	"config": {
		"db_type": "MYSQL",
		"db_username": "username",
		"db_password": "password",
		"db_host": "host",
		"db_port": "port",
		"db_name": "name"
	},
	"models": [{
			"name": "User",
			"fields": [{
					"name": "username",
					"type": "string",
					"required": true,
					"default": "asd",
					"unique": true
				},
				{
					"name": "password",
					"type": "string",
					"required": false,
					"default": "",
					"unique": false
				}
			]
		},
		{
			"name": "Product",
			"fields": [{
					"name": "name",
					"type": "string",
					"required": true,
					"default": "asd",
					"unique": true
				},
				{
					"name": "quantity",
					"type": "int",
					"required": false,
					"default": "3",
					"unique": false
				}
			]
		}
	]
}
*/
