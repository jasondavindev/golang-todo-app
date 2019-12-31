package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/common"
)

type TaskValidator struct {
	Name        string `json:"name" binding:"required,min=10,max=100"`
	Description string `json:"description" binding:"omitempty"`
	UserID      uint   `json:"user_id" binding:"required"`
	TaskModel   Task   `json:"-"`
}

func NewTaskValidator() TaskValidator {
	return TaskValidator{}
}

func (self *TaskValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}

	self.TaskModel = Task{
		Name:        self.Name,
		Description: self.Description,
		UserID:      self.UserID,
	}

	return nil
}
