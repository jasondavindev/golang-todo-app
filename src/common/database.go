package common

import (
	"os"

	"github.com/jasondavindev/golang-todo-app/src/configs"
	"github.com/jinzhu/gorm"
)

var instance *gorm.DB

func InitDB() *gorm.DB {
	dbDriver := configs.GetConfig().Database.Driver
	dbURL := configs.GetConfig().Database.URL
	db, err := gorm.Open(dbDriver, dbURL)

	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	instance = db

	if os.Getenv("GOENV") == "development" {
		instance = instance.Debug()
	}

	return instance
}

func GetDB() *gorm.DB {
	return instance
}
