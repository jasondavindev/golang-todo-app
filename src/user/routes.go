package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(r *gin.RouterGroup) {
	r.POST("/", UserRegistration)
	r.GET("/:name/tasks", UserTasks)
	r.GET("/:name", UserFind)
}

func UserFind(c *gin.Context) {
	userSearch := c.Param("name")

	user, err := FindOne(&User{Name: userSearch})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.Response()})
}

func UserRegistration(c *gin.Context) {
	userValidator := NewUserValidator()

	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userValidator.UserModel.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userValidator.UserModel.Response()})
}

func UserTasks(c *gin.Context) {
	name := c.Param("name")
	user, err := FindOne(&User{Name: name})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks := user.GetTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
