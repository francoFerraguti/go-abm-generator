package templates

func MainGo() string {
	return `package main

&&IMPORTS&&

func main() {
	dbhandler.ConnectToDatabase()
	router.ConfigureRouter()
	router.CreateRouter()
	router.RunRouter()
}
`
}

func MiddlewareGo() string {
	return `package middleware

&&IMPORTS&&

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token := authentication.GetTokenData(tokenString)

		if token.&&USERNAME_FIELD_UPPER&& == "" || tokenString == "" {
		    c.JSON(401, "Authentication error")
	    	c.Abort()
			return
		}

		c.Set("&&USERNAME_FIELD_LOWER&&", token.&&USERNAME_FIELD_UPPER&&)
	}
}
`
}

func ConfigGo() string {
	return `package config

type Config struct {
	ENV			string
	PORT 		string
	&&CONFIG_AUTHENTICATION_FIELD&&
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
		PORT:			&&CONFIG_PORT&&,
		&&CONFIG_AUTHENTICATION_VALUE&&,
		DB_TYPE:       	&&CONFIG_DB_TYPE&&,
		DB_USERNAME:    &&CONFIG_DB_USERNAME&&,
		DB_PASSWORD:    &&CONFIG_DB_PASSWORD&&,
		DB_HOST:      	&&CONFIG_DB_HOST&&,
		DB_PORT:       	&&CONFIG_DB_PORT&&,
		DB_NAME:       	&&CONFIG_DB_NAME&&,
	}
}
`
}

func RouterGo() string {
	return `package router

&&IMPORTS&&

var router *gin.Engine

func ConfigureRouter() {
	if config.GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

	&&AUTHENTICATION_ENDPOINTS&&

	api := router.Group("/"&&GIN_MIDDLEWARE_STRING&&)
	{
		&&ENDPOINTS&&
    }
}

func RunRouter() {
	router.Run(":" + config.GetConfig().PORT)
}
`
}
