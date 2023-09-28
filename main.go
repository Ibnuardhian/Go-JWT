package main

import (
	"log"

	"github.com/Ibnuardhian/go-jwt/controller"
	"github.com/Ibnuardhian/go-jwt/database"
	"github.com/Ibnuardhian/go-jwt/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env Files")
	}
	r := gin.Default()
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.ReqAuth, controller.Validate)

	r.Run()
}
