package common

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dbURL)

	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
