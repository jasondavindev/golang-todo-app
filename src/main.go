package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/common"
	"github.com/jasondavindev/golang-todo-app/src/task"
	"github.com/jasondavindev/golang-todo-app/src/user"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := common.Init()
	Migrate()
	defer db.Close()

	router := gin.Default()

	api := router.Group("/api")

	user.UserRegister(api.Group("/users"))
	task.TaskRegister(api.Group("/tasks"))

	router.Run(":3000")
}

func Migrate() {
	user.AutoMigrate()
	task.AutoMigrate()
}
