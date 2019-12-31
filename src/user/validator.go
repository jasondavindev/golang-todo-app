package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/common"
)

type UserValidator struct {
	Name      string `json:"name" binding:"required,alphanum,min=4,max=25"`
	Password  string `json:"password" binding:"required,min=8,max=30"`
	Email     string `json:"email" binding:"required,email"`
	UserModel *User  `json:"-"`
}

func (self *UserValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}

	self.UserModel = &User{
		Email:    self.Email,
		Name:     self.Name,
		Password: self.Password,
	}

	return nil
}

func NewUserValidator() UserValidator {
	return UserValidator{}
}
