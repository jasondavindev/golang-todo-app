package user

import (
	"errors"

	"github.com/jasondavindev/golang-todo-app/src/common"
	"github.com/jasondavindev/golang-todo-app/src/task"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string      `json:"name" gorm:"unique_index;not null"`
	Password string      `json:"password"`
	Tasks    []task.Task `json:"-"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&User{})
}

func FindOne(cond interface{}) (User, error) {
	var user User
	err := common.GetDB().First(&user, cond).Error
	return user, err
}

func FindAll(cond ...interface{}) ([]User, error) {
	var users []User
	err := common.GetDB().Find(&users, cond...).Error
	return users, err
}

func Save(user *User) error {
	err := common.GetDB().Save(user).Error
	return err
}

func (user *User) tasks() []task.Task {
	db := common.GetDB()
	var tasks []task.Task
	db.Model(&user).Related(&tasks)
	return tasks
}

func (user *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return nil
}
