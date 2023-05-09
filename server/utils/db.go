package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gsCheck/model"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(model.Dept{})
	DB.AutoMigrate(model.User{})
}
