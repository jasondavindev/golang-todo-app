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

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserRegistration(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := (&user).Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
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
