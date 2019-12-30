package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(r *gin.RouterGroup) {
	r.POST("/", UserRegistration)
	r.GET("/:name", UserRetrieve)
}

func UserRetrieve(c *gin.Context) {
	userSearch := c.Param("name")

	user, err := FindOne(&User{Name: userSearch})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserRegistration(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.setPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Save(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
