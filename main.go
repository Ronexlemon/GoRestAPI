package main

import (
	//"example/restapi/album"
	"example/restapi/database"
	"example/restapi/model"
)

func main() {
	loadDatabase()
}
func loadDatabase(){
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

// func main(){
// 	Generic()
// 	GenericCall()
// }
