package main

import (
	"./src/service"
	"fmt"
	"model"
)

var APPNAME = "RayBook"

func main(){
	fmt.Sprintf("this project is %v", APPNAME)
	service.StartWebServer("127.0.0.1","8080")
	initializeBoltClient()                 // NEW

}


// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	model.DBClient = &model.BoltClient{}
	model.DBClient.OpenBoltDb()
	model.DBClient.Seed()
}





