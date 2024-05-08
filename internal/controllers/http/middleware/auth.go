package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"task-manager/internal/controllers/http/auth"
)

func CheckTokenMiddleware(tokenService auth.IJWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		splitToken := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization required"})
			return
		}

		token := splitToken[1]

		parsedToken, err := tokenService.ParseAccessToken(token)
		if err != nil {
			logrus.Error("auth middleware :", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token or token is expired"})
			return
		}

		fmt.Println(parsedToken.Claims)

		c.Next()
	}
}
