package task

import (
	"github.com/jasondavindev/golang-todo-app/src/common"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	UserID      uint   `json:"user_id"`
	Done        bool   `json:"done"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Task{})
}

func Save(task *Task) error {
	db := common.GetDB()
	return db.Save(task).Error
}

func FindOne(cond ...interface{}) (Task, error) {
	db := common.GetDB()
	var task Task
	err := db.First(&task, cond...).Error
	return task, err
}
