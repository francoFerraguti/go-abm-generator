package main

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/liteByte/frango"
	"strconv"
	"strings"
)

func getFileControllerAuthenticationGo(projectName string, models []structs.ModelStruct) string {
	authenticationModel := structs.ModelStruct{}
	usernameField := structs.FieldStruct{}
	passwordField := structs.FieldStruct{}

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

func getFileAuthenticationGo(projectName string, models []structs.ModelStruct) string {
	usernameField := structs.FieldStruct{}

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

func getFileDBHandlerGo(projectPath string, models []structs.ModelStruct) string {
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

func getFileDBSchemaGo(models []structs.ModelStruct) string {
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

func getFileModelGo(projectName string, needAuthentication bool, model structs.ModelStruct) string {
	getByString := ""
	deleteByString := ""
	updateByString := ""
	checkLoginString := ""
	authenticationUsername := structs.FieldStruct{}
	authenticationPassword := structs.FieldStruct{}
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

func getFileControllerGo(projectName string, model structs.ModelStruct) string {
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

func getFileDocumentation(needAuthentication bool, models []structs.ModelStruct) string {
	s := ""

	if needAuthentication {

		authModel := structs.ModelStruct{}
		usernameField := structs.FieldStruct{}
		passwordField := structs.FieldStruct{}

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

func getEndpointBody(model structs.ModelStruct, onlySend bool) string {
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
