package infrastructures

import (
	"net/http"
	"strings"

	"github.com/abe16s/Online-Marketplace-API/usecases/interfaces"
	"github.com/gin-gonic/gin"
)

func RefreshMiddleware(jwtservice interfaces.IJwtService) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		email, isSeller, err := jwtservice.ValidateToken(authParts[1], true)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("refresh_token", map[string]interface{}{"email": email, "isSeller": isSeller})	
		c.Next()
	}
}