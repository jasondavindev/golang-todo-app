package user

import "github.com/gin-gonic/gin"

import "net/http"

import "github.com/dgrijalva/jwt-go"

import "github.com/jasondavindev/golang-todo-app/src/common"

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
			return common.GetJwtKey(), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
