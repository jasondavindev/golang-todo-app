package task

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"
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
	UUID        string `json:"uuid" gorm:"unique"`
}

type TaskRetrieve struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Done        bool   `json:"done"`
	UUID        string `json:"uuid"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Task{})
}

func (task *Task) BeforeCreate() {
	uuid, err := uuid.NewUUID()

	if err != nil {
		return
	}

	task.UUID = uuid.String()
	task.Slug = slug.Make(task.Name)
}

func (task *Task) Save() error {
	db := common.GetDB()
	return db.Save(task).Error
}

func FindOne(cond ...interface{}) (Task, error) {
	db := common.GetDB()
	var task Task
	err := db.First(&task, cond...).Error
	return task, err
}

func (task *Task) Update(data Task) error {
	db := common.GetDB()
	err := db.Model(task).Update(data).Error
	return err
}

func FindBy(cond ...Task) ([]Task, error) {
	db := common.GetDB()
	var tasks []Task
	err := db.Find(&tasks, cond).Error
	return tasks, err
}
