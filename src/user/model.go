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
	Password string      `json:"password" gorm:"not null"`
	Email    string      `json:"email" gorm:"unique,not null"`
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

func (user *User) Save() error {
	err := common.GetDB().Save(user).Error
	return err
}

func (user *User) Update(data User) error {
	return common.GetDB().Model(user).Update(data).Error
}

func (user *User) GetTasks() []task.TaskRetrieve {
	db := common.GetDB()
	var tasks []task.TaskRetrieve
	db.Find(&[]task.Task{}, &task.Task{UserID: user.ID}).Scan(&tasks)
	return tasks
}

func (user *User) BeforeSave() {
	if err := user.setPassword(); err != nil {
		panic(err.Error())
	}
}

func (user *User) setPassword() error {
	if len(user.Password) == 0 {
		return errors.New("password should not be empty")
	}

	bytePassword := []byte(user.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return nil
}
