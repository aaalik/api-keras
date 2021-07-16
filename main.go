package main

import (
	"github.com/aaalik/api-keras/bootstrap"
	"github.com/aaalik/api-keras/helper"
	modelItem "github.com/aaalik/api-keras/model/item"
	modelUser "github.com/aaalik/api-keras/model/user"
	"github.com/aaalik/api-keras/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bootstrap.InitLogger()
	helper.Log = bootstrap.Log

	bootstrap.Connect()

	modelItem.Migrate(bootstrap.DB)
	modelUser.Migrate(bootstrap.DB)

	routers.SetupRouter()

	forever := make(chan bool)
	<-forever
}
