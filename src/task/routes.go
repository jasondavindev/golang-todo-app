package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskRegister(r *gin.RouterGroup) {
	r.POST("/:uuid/done", TaskDone)
	r.POST("/", TaskCreation)
	r.GET("/:slug", Retrieve)
}

func TaskCreation(c *gin.Context) {
	var task Task

	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := (&task).Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func Retrieve(c *gin.Context) {
	slug := c.Param("slug")

	task, err := FindOne(&Task{Slug: slug})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func TaskDone(c *gin.Context) {
	uuid := c.Param("uuid")

	task, err := FindOne(&Task{UUID: uuid})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if task.Done {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
		return
	}

	if err = task.Update(Task{Done: true}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}
