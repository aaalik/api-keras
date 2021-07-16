package bootstrap

import (
	"fmt"
	"os"

	"github.com/aaalik/api-keras/helper"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() {
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DATABASE_HOST := os.Getenv("DATABASE_HOST")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=true&loc=Local", DATABASE_USER, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_NAME)

	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		helper.Log.Error(err)
	}

	DB = database
}
