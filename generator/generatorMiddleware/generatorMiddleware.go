package generatorMiddleware

import (
	"github.com/liteByte/frango"
)

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