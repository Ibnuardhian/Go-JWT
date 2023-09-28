package database

import (
	"log"
	"os"

	"github.com/Ibnuardhian/go-jwt/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("dsn")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to Database")
	} else {
		log.Println("Connect Successfully")
	}
	DB = database
	database.AutoMigrate(&models.User{})
}
