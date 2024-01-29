package database

import (
	"fmt"
	//"gorm.io/driver/postgres"
	"os"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	//var err error
	err := godotenv.Load()
    if err != nil {
        // Handle error loading .env file
		panic(err)
    }
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s  user=%s  password =%s dbname=%s port =%s sslmode=disable ", host, username, password, databaseName, port)
	fmt.Println(dsn)
}
