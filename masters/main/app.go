package main

import (
	"github.com/inact25/E-WarungApi/configs"
	"github.com/inact25/E-WarungApi/masters/apis"
	"github.com/inact25/E-WarungApi/utils"
)

func main() {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)
	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	apis.Init(myRoute, db)
	configs.RunServer(myRoute)
}
