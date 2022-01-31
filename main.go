package main

import (
	auth "crowdfunding/Auth"
	handler "crowdfunding/Handler"
	users "crowdfunding/Users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "jun:jun123@tcp(127.0.0.1:3306)/gopractice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// repo
	userRepository := users.NewRepository(db)

	// service
	userService := users.NewService(userRepository)
	authService := auth.NewService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/checkEmail", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)
	router.Run()
}
