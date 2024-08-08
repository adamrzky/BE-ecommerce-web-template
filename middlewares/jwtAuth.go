package middlewares

import (
	"BE-ecommerce-web-template/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			var statusCode int
			var message string

			switch err.Error() {
			case "token contains an invalid number of segments":
				statusCode = http.StatusUnauthorized
				message = "Unauthorized: Token is malformed"
			case "token is expired":
				statusCode = http.StatusUnauthorized
				message = "Unauthorized: Token has expired"
			case "token not valid yet":
				statusCode = http.StatusUnauthorized
				message = "Unauthorized: Token is not valid yet"
			default:
				statusCode = http.StatusUnauthorized
				message = "Unauthorized: " + err.Error()
			}

			c.JSON(statusCode, gin.H{
				"status":  "error",
				"message": message,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
