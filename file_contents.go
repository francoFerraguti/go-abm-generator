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

func getFileControllerAuthenticationGo(projectName string, models []ModelStruct) string {
	authenticationModel := ModelStruct{}
	usernameField := FieldStruct{}
	passwordField := FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				authenticationModel = model
				usernameField = field
			}

			if field.AuthenticationPassword {
				authenticationModel = model
				passwordField = field
			}
		}
	}

	return `package authentication

import (
	"` + projectName + `/structs"
	"` + projectName + `/authentication"
	"` + projectName + `/models/` + frango.FirstLetterToLower(authenticationModel.Name) + `"
    "github.com/liteByte/frango"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginStruct, err := structs.ParseBodyIntoLoginStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	loginStruct.` + frango.FirstLetterToUpper(passwordField.Name) + ` = frango.Hash(loginStruct.` + frango.FirstLetterToUpper(usernameField.Name) + `, loginStruct.` + frango.FirstLetterToUpper(passwordField.Name) + `)

	if err = ` + frango.FirstLetterToLower(authenticationModel.Name) + `.CheckLogin(loginStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	token := authentication.CreateToken(loginStruct.` + frango.FirstLetterToUpper(usernameField.Name) + `)

	c.JSON(200, token)
}

func Signup(c *gin.Context) {
	` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct, err := structs.ParseBodyInto` + frango.FirstLetterToUpper(authenticationModel.Name) + `Struct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(passwordField.Name) + ` = frango.Hash(` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(usernameField.Name) + `, ` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(passwordField.Name) + `)

	if err = ` + frango.FirstLetterToLower(authenticationModel.Name) + `.Create(` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct); err != nil {
		c.JSON(500, err.Error())
		return		
	}

	c.JSON(200, "Signup successful")
}`
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
	usernameField := FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				usernameField = field
			}
		}
	}

	usernameInStruct := frango.FirstLetterToUpper(usernameField.Name) + " " + usernameField.Type

	return `package authentication

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"` + projectName + `/config"
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

func getFileStructsGo(projectName string, needAuthentication bool, models []ModelStruct) string {
	structsString := ""
	authStructString := ""

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

	if needAuthentication {
		authStructString += "type LoginStruct struct {\n"

		for _, model := range models {
			for _, field := range model.Fields {
				if field.AuthenticationUsername || field.AuthenticationPassword {
					authStructString += "	" + frango.FirstLetterToUpper(field.Name) + " " + field.Type + "\n"
				}
			}
		}

		authStructString += "}\n\n"

		authStructString += "func ParseBodyIntoLoginStruct(body io.ReadCloser) (LoginStruct, error) {\n"
		authStructString += "    bodyBytes, _ := ioutil.ReadAll(body)\n"
		authStructString += "    loginStruct := LoginStruct{}\n"
		authStructString += "    err := json.Unmarshal(bodyBytes, &loginStruct)\n"
		authStructString += "    return loginStruct, err\n"
		authStructString += "}\n\n"
	}

	return `package structs

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

` + structsString + authStructString
}

func getFileRouterGo(projectName string, needAuthentication bool, models []ModelStruct) string {
	importString := ""
	endpointsString := ""
	authMiddlewareString := ""
	authEndpoints := ""
	authImports := ""

	if needAuthentication {
		authMiddlewareString = ", middleware.ValidateToken()"

		authEndpoints = `	public := router.Group("/")
	{
		public.POST("/signup", authentication.Signup)
		public.POST("/login", authentication.Login)
	}`

		authImports = `	"` + projectName + `/middleware"
	"` + projectName + `/controllers/authentication"`
	}

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
` + importString + authImports + `
)

var router *gin.Engine

func ConfigureRouter() {
	if config.GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

` + authEndpoints + `

	api := router.Group("/"` + authMiddlewareString + `)
	{
` + endpointsString + `    }
}

func RunRouter() {
	router.Run(":" + config.GetConfig().PORT)
}
`
}

func getFileModelGo(projectName string, needAuthentication bool, model ModelStruct) string {
	getByString := ""
	deleteByString := ""
	updateByString := ""
	checkLoginString := ""
	authenticationUsername := FieldStruct{}
	authenticationPassword := FieldStruct{}
	authImportString := ""

	for _, field := range model.Fields {
		if field.AuthenticationUsername {
			authenticationUsername = field
		}
		if field.AuthenticationPassword {
			authenticationPassword = field
		}

		if !field.Unique {
			continue
		}

		updateByString += modelUpdateBy(model, field)
		getByString += modelGetBy(model, field)
		deleteByString += modelDeleteBy(model, field)
	}

	if needAuthentication && authenticationUsername.Name != "" && authenticationPassword.Name != "" {
		checkLoginString += "func CheckLogin(loginStruct structs.LoginStruct) error {\n"
		checkLoginString += "	var exists bool\n\n"
		checkLoginString += "	err := dbhandler.GetDatabase().QueryRow(`SELECT EXISTS (SELECT 1 FROM " + frango.FirstLetterToUpper(model.Name) + " WHERE " + authenticationUsername.Name + " = ? AND " + authenticationPassword.Name + " = ? LIMIT 1)`, loginStruct." + frango.FirstLetterToUpper(authenticationUsername.Name) + ", loginStruct." + frango.FirstLetterToUpper(authenticationPassword.Name) + ").Scan(&exists)\n"
		checkLoginString += "	if err != nil {\n"
		checkLoginString += "		return err\n"
		checkLoginString += "	}\n\n"
		checkLoginString += "	if !exists {\n"
		checkLoginString += "		return errors.New(`Login failed`)\n"
		checkLoginString += "	}\n\n"
		checkLoginString += "	return nil\n"
		checkLoginString += "}\n\n"

		authImportString = `"errors"`
	}

	return `package ` + strings.ToLower(model.Name) + `

import (
	` + authImportString + `
	"` + projectName + `/dbhandler"
    "` + projectName + `/structs"
)

` + checkLoginString + modelCreate(model) + updateByString + getByString + deleteByString
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

func getFileDocumentation(needAuthentication bool, models []ModelStruct) string {
	s := ""

	if needAuthentication {

		authModel := ModelStruct{}
		usernameField := FieldStruct{}
		passwordField := FieldStruct{}

		for _, model := range models {
			for _, field := range model.Fields {
				if field.AuthenticationUsername {
					usernameField = field
					authModel = model
				}
				if field.AuthenticationPassword {
					passwordField = field
					authModel = model
				}
			}
		}

		title := "Signup"
		description := "Creates a user on the database"
		endpoint := "/signup"
		method := "POST"
		url_params := ""
		headers := ""
		body := getEndpointBody(authModel, true)
		code200 := "Signup successful"
		code400 := "Request body is wrong"
		code401 := ""
		code500 := "Username already in use"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

		title = "Login"
		description = "Returns an access token for the provided user"
		endpoint = "/login"
		method = "POST"
		url_params = ""
		headers = ""
		body = `{
	"` + usernameField.Name + `": "` + usernameField.Type + `",
	"` + passwordField.Name + `": "` + passwordField.Type + `"
}`
		code200 = "a63ab36162a4f4ee6622ccd787b0a048c26b93acfc05c6b1843659b253c3c00b //authentication token"
		code400 = "Request body is wrong"
		code401 = ""
		code500 = "Wrong username or password"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)
	}

	for _, model := range models {
		//Create
		title := "Create " + model.Name
		description := "Creates a " + frango.FirstLetterToLower(model.Name) + " on the database"
		endpoint := "/" + frango.FirstLetterToLower(model.Name)
		method := "POST"
		url_params := ""
		headers := "Authorization: Token"
		body := getEndpointBody(model, true)
		code200 := model.Name + " created successfully"
		code400 := "Request body is wrong"
		code401 := "Unauthorized"
		code500 := "Server error"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

		for _, field := range model.Fields {
			if !field.Unique {
				continue
			}

			//Get
			title = "Get " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Returns a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "GET"
			url_params = ""
			headers = "Authorization: Token"
			body = ""
			code200 = getEndpointBody(model, false)
			code400 = ""
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

			//Update
			title = "Update " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Updates a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "PUT"
			url_params = ""
			headers = "Authorization: Token"
			body = getEndpointBody(model, true)
			code200 = model.Name + " updated successfully"
			code400 = "Request body is wrong"
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

			//Delete
			title = "Delete " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Deletes a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "DELETE"
			url_params = ""
			headers = "Authorization: Token"
			body = ""
			code200 = model.Name + " deleted successfully"
			code400 = ""
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)
		}
	}

	return s
}

func getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500 string) string {
	if url_params == "" {
		url_params = "-"
	}
	if headers == "" {
		headers = "-"
	}
	if body == "" {
		body = "-"
	}
	if code400 == "" {
		code400 = "-"
	}
	if code401 == "" {
		code401 = "-"
	}

	return `**` + title + `**
` + description + `

**` + endpoint + `** [` + method + `]

**URL Parameters**
` + url_params + `

**Headers**
` + headers + `

**Request Body**
` + body + `

**Success 200 Response**
` + code200 + `

**Bad Request 400 Response**
` + code400 + `

**Unauthorized 401 Response**
` + code401 + `

**Internal Server Error 500 Response**
` + code500 + `

-----------------------------------------------------
`
}

func getEndpointBody(model ModelStruct, onlySend bool) string {
	s := "{\n"

	for i, field := range model.Fields {

		if field.AutoGenerated && onlySend {
			continue
		}

		fieldTypeString := ""

		if field.Type == "string" {
			fieldTypeString = "`string`"
		} else {
			fieldTypeString = field.Type
		}

		if i != len(model.Fields)-1 {
			fieldTypeString += ","
		}

		s += "	`" + frango.FirstLetterToLower(field.Name) + "`: " + fieldTypeString + "\n"
	}

	s += "}"

	return s
}
