package user

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/configs"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.AbortWithStatus(http.StatusBadRequest)
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetConfig().JWT.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
