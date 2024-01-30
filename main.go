package main

import (
	//"example/restapi/album"
	"example/restapi/database"
	"example/restapi/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
}
func loadDatabase(){
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}
func loadEnv(){
	err:= godotenv.Load(".env.local")
	if err !=nil{
		log.Fatal("Error loading .env file")
	}

}

// func main(){
// 	Generic()
// 	GenericCall()
// }
