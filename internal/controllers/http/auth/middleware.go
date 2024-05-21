package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func CheckTokenMiddleware(tokenService jwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		splitToken := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(splitToken) != 2 {
			logrus.Error("CheckTokenMiddleware | token is wrong or not provider or wrong of format bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization required"})
			return
		}

		token := splitToken[1]

		claims, err := tokenService.ParseAccessToken(token)
		if err != nil {
			logrus.Error("CheckTokenMiddleware :", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token or token is expired"})
			return
		}

		fmt.Println(claims)

		c.Next()
	}
}
