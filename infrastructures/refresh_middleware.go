package infrastructures

import (
	"net/http"

	"github.com/abe16s/Online-Marketplace-API/usecases/interfaces"
	"github.com/gin-gonic/gin"
)

func RefreshMiddleware(jwtservice interfaces.IJwtService) gin.HandlerFunc {

	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
            c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "refresh token is required"})
            c.Abort()
            return
        }

		email, isSeller, err := jwtservice.ValidateToken(refreshToken, true)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("refresh_token", map[string]interface{}{"email": email, "isSeller": isSeller})	
		c.Next()
	}
}