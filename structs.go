package main

type Config struct {
	DB_TYPE     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

type Model struct {
	Name string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Required bool
	Default string
}

func getModelsArray(modelsInterface []interface{}) []Model {
	models := make([]Model, 0)

	for key, _ := range modelsInterface {
		modelInterface := modelsInterface[key].(map[string]interface {})
		fieldsInterface := modelInterface["fields"].([]interface{})

		model := Model{}
		model.Name = modelInterface["name"].(string)

		fields := make([]Field, 0)

		for key2, _ := range fieldsInterface {
			fieldInterface := fieldsInterface[key2].(map[string]interface {})

			field := Field{}
			field.Name = fieldInterface["name"].(string)
			field.Type = fieldInterface["name"].(string)
			field.Required = fieldInterface["name"].(bool)
			field.Default = fieldInterface["name"].(string)

			fields = append(fields, field)
		}

		model.Fields = fields
		models = append(models, model)
	}

	return models
}

func getConfigStruct(config map[string]interface{}) Config {
	return Config {
		DB_TYPE: config["db_type"].(string),
		DB_USERNAME: config["db_username"].(string),
		DB_PASSWORD: config["db_password"].(string),
		DB_HOST: config["db_host"].(string),
		DB_PORT: config["db_port"].(string),
		DB_NAME: config["db_name"].(string),
	}
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
					"default": "asd"
				},
				{
					"name": "password",
					"type": "string",
					"required": false,
					"default": ""
				}
			]
		},
		{
			"name": "Product",
			"fields": [{
					"name": "name",
					"type": "string",
					"required": true,
					"default": "asd"
				},
				{
					"name": "quantity",
					"type": "int",
					"required": false,
					"default": 3
				}
			]
		}
	]
}
*/