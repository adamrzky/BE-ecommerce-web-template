package middlewares

import (
	"BE-ecommerce-web-template/utils/resp"
	"BE-ecommerce-web-template/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := token.TokenValid(c)
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

		if !hasValidRole(claims, roles) {
			resp.NewResponseForbidden(c, "forbidden")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

// -----------------------------------------------------------------------------
// Helper func
// -----------------------------------------------------------------------------

func hasValidRole(claims *token.JWTClaims, roles []string) bool {
	if len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if claims.Role == role {
			return true
		}
	}

	return false
}
