package main

import (
	"github.com/liteByte/frango"
	"strconv"
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

func getFileMiddlewareGo(projectName string, models []ModelStruct) string {
	usernameField := FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				usernameField = field
			}
		}
	}

	return `package middleware

import (
	"github.com/gin-gonic/gin"
	"` + projectName + `/authentication"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token := authentication.GetTokenData(tokenString)

		if token.` + frango.FirstLetterToUpper(usernameField.Name) + ` == "" || tokenString == "" {
		    c.JSON(401, "Authentication error")
	    	c.Abort()
			return
		}

		c.Set("` + frango.FirstLetterToLower(usernameField.Name) + `", token.` + frango.FirstLetterToUpper(usernameField.Name) + `)
	}
}
`
}

func getFileAuthenticationGo(projectName string, models []ModelStruct) string {
	//authenticationModel := ModelStruct{}
	usernameField := FieldStruct{}
	//passwordField := FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				usernameField = field
				//authenticationModel = model
			}
			if field.AuthenticationPassword {
				//passwordField = field
				//authenticationModel = model
			}
		}
	}

	usernameInStruct := frango.FirstLetterToUpper(usernameField.Name) + " " + usernameField.Type

	return `package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	` + usernameInStruct + `
	jwt.StandardClaims
}

type Token struct {
	` + usernameInStruct + `
}

func CreateToken(` + frango.FirstLetterToLower(usernameField.Name) + " " + usernameField.Type + `) string {
    claims := CustomClaims {
        ` + frango.FirstLetterToLower(usernameField.Name) + `,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
            Issuer:    "TODO change to application name",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, _ := token.SignedString([]byte(config.GetConfig().JWT_SECRET))

    return tokenString
}

func GetTokenData(tokenString string) Token {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT_SECRET), nil
	})
	if err != nil {
		return Token{}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return Token{
			claims.` + frango.FirstLetterToUpper(usernameField.Name) + `,
			claims.StandardClaims.ExpiresAt, 
			claims.StandardClaims.IssuedAt, 
			claims.StandardClaims.Issuer,
		}
	} else {
		return Token{}
	}
}

`
}

func getFileConfigGo(projectName string, needAuthentication bool, config ConfigStruct) string {
	authenticationString := ""
	authenticationString2 := ""
	if needAuthentication {
		authenticationString = "	JWT_SECRET	string"
		authenticationString2 = `		JWT_SECRET:		"` + frango.GetRandomString(32) + `",`
	}

	return `package config

type Config struct {
	ENV			string
	PORT 		string
` + authenticationString + `
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
		ENV:			"develop",
		PORT:			"` + config.Port + `",
` + authenticationString2 + `
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

func getFileDBHandlerGo(projectPath string, models []ModelStruct) string {
	schemaString := "func createSchema() {\n"

	for _, model := range models {
		schemaString += "	create" + model.Name + "Table()\n"
	}

	schemaString += "}"

	return `package dbhandler

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/liteByte/frango"
    "` + projectPath + `/config"
)

var db *sql.DB

func ConnectToDatabase() {
    var err error
	
	db, err = sql.Open(config.GetConfig().DB_TYPE, config.GetConfig().DB_USERNAME + ":" + config.GetConfig().DB_PASSWORD + "@tcp(" + config.GetConfig().DB_HOST + ":" + config.GetConfig().DB_PORT + ")/" + config.GetConfig().DB_NAME)
	frango.PrintErr(err)
    
    err = db.Ping()
    frango.PrintErr(err)

    createSchema()
}

func GetDatabase() *sql.DB {
    return db
}

` + schemaString
}

func getFileDBSchemaGo(models []ModelStruct) string {
	schemaString := ""

	for _, model := range models {
		schemaString += "func create" + model.Name + "Table() {\n"
		schemaString += "	query := `CREATE TABLE IF NOT EXISTS " + model.Name + " (\n"

		for i, field := range model.Fields {
			fieldString := field.Name

			max := "255"
			if field.Max != 0 {
				max = strconv.Itoa(field.Max)
			}

			if field.Type == "string" {
				fieldString += " varchar(" + max + ")"
			} else if field.Type == "bool" {
				fieldString += " tinyint"
			} else {
				fieldString += " " + field.Type
			}

			if field.Unique {
				fieldString += " UNIQUE"
			}

			if field.Required {
				fieldString += " NOT NULL"
			}

			if field.Default != "" {
				if field.Type == "bool" {
					if field.Default == "true" {
						fieldString += " DEFAULT '1'"
					} else if field.Default == "false" {
						fieldString += " DEFAULT '0'"
					}
				} else {
					fieldString += " DEFAULT '" + field.Default + "'"
				}
			}

			if field.AutoGenerated && field.Type == "int" {
				fieldString += " AUTO_INCREMENT"
			}

			if i != len(model.Fields)-1 {
				fieldString += ","
			}

			fieldString += "\n"
			schemaString += "		" + fieldString
		}
		schemaString += "	);`\n"
		schemaString += "	_, err := db.Exec(query)\n"
		schemaString += "	frango.PrintErr(err)\n"
		schemaString += "}\n\n"
	}

	return `package dbhandler

import (
    "github.com/liteByte/frango"
)

` + schemaString
}

func getFileStructsGo(projectName string, models []ModelStruct) string {
	structsString := ""
	for _, model := range models {
		structsString += "type " + model.Name + "Struct struct {\n"
		for _, field := range model.Fields {
			structsString += "	" + frango.FirstLetterToUpper(field.Name) + " " + field.Type + "\n"
		}
		structsString += "}\n\n"

		structsString += "func ParseBodyInto" + model.Name + "Struct(body io.ReadCloser) (" + model.Name + "Struct, error) {\n"
		structsString += "    bodyBytes, _ := ioutil.ReadAll(body)\n"
		structsString += "    " + frango.FirstLetterToLower(model.Name) + "Struct := " + model.Name + "Struct{}\n"
		structsString += "    err := json.Unmarshal(bodyBytes, &" + frango.FirstLetterToLower(model.Name) + "Struct)\n"
		structsString += "    return " + frango.FirstLetterToLower(model.Name) + "Struct, err\n"
		structsString += "}\n\n"
	}

	return `package structs

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

` + structsString
}

func getFileRouterGo(projectName string, models []ModelStruct) string {
	importString := ""
	endpointsString := ""

	for _, model := range models {
		importString += "	`" + projectName + "/controllers/" + frango.FirstLetterToLower(model.Name) + "`\n"

		endpointsString += "		api.POST(`/" + frango.FirstLetterToLower(model.Name) + "`, " + frango.FirstLetterToLower(model.Name) + ".Create)\n"

		for _, field := range model.Fields {
			if !field.Unique {
				continue
			}

			endpointsString += "		api.GET(`/" + frango.FirstLetterToLower(model.Name) + "/" + frango.FirstLetterToLower(field.Name) + "/:" + frango.FirstLetterToLower(field.Name) + "`, " + frango.FirstLetterToLower(model.Name) + ".GetBy" + frango.FirstLetterToUpper(field.Name) + ")\n"
			endpointsString += "		api.PUT(`/" + frango.FirstLetterToLower(model.Name) + "/" + frango.FirstLetterToLower(field.Name) + "/:" + frango.FirstLetterToLower(field.Name) + "`, " + frango.FirstLetterToLower(model.Name) + ".UpdateBy" + frango.FirstLetterToUpper(field.Name) + ")\n"
			endpointsString += "		api.DELETE(`/" + frango.FirstLetterToLower(model.Name) + "/" + frango.FirstLetterToLower(field.Name) + "/:" + frango.FirstLetterToLower(field.Name) + "`, " + frango.FirstLetterToLower(model.Name) + ".DeleteBy" + frango.FirstLetterToUpper(field.Name) + ")\n"
		}
	}

	return `package router

import (
	"github.com/gin-gonic/gin"
	"` + projectName + `/config"
` + importString + `)

var router *gin.Engine

func ConfigureRouter() {
	if config.GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

	api := router.Group("/")
	{
` + endpointsString + `    }
}

func RunRouter() {
	router.Run(":" + config.GetConfig().PORT)
}
`
}

func getFileModelGo(projectName string, model ModelStruct) string {
	getByString := ""
	deleteByString := ""
	updateByString := ""

	for _, field := range model.Fields {
		if !field.Unique {
			continue
		}

		updateByString += modelUpdateBy(model, field)
		getByString += modelGetBy(model, field)
		deleteByString += modelDeleteBy(model, field)
	}

	return `package ` + strings.ToLower(model.Name) + `

import (
	"` + projectName + `/dbhandler"
    "` + projectName + `/structs"
)

` + modelCreate(model) + updateByString + getByString + deleteByString
}

func getFileControllerGo(projectName string, model ModelStruct) string {
	getByString := ""
	deleteByString := ""
	updateByString := ""
	frangoImport := ""

	for _, field := range model.Fields {
		if field.Type != "string" && field.Unique {
			frangoImport = "	\n`github.com/liteByte/frango`\n"
		}

		if !field.Unique {
			continue
		}

		updateByString += controllerUpdateBy(model, field)
		getByString += controllerGetBy(model, field)
		deleteByString += controllerDeleteBy(model, field)
	}

	return `package ` + strings.ToLower(model.Name) + `

import (
	"encoding/json"` + frangoImport + `
	"` + projectName + `/models/` + strings.ToLower(model.Name) + `"
    "` + projectName + `/structs"
	"github.com/gin-gonic/gin"
)

` + controllerCreate(model) + updateByString + getByString + deleteByString
}
