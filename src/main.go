package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/common"
	"github.com/jasondavindev/golang-todo-app/src/configs"
	"github.com/jasondavindev/golang-todo-app/src/task"
	"github.com/jasondavindev/golang-todo-app/src/user"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := configs.Init(); err != nil {
		panic(err)
	}

	db := common.InitDB()
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
