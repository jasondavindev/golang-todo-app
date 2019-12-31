package user

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
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

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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

func (user *User) GetTasks() []task.TaskResponse {
	db := common.GetDB()
	var tasks []task.Task
	db.Find(&tasks, &task.Task{UserID: user.ID})
	return task.Map(tasks)
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
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	return nil
}

func (user *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
