package main

import (
	handler "crowdfunding/Handler/Users"
	auth "crowdfunding/Services/Auth"
	users "crowdfunding/Services/Users"
	"crowdfunding/database"
	"crowdfunding/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.NewConnection()

	// repo
	userRepository := users.NewRepository(db)

	// service
	userService := users.NewService(userRepository)
	authService := auth.NewService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)

	// middleware
	// authMiddleware := middleware.Middleware(authService, userService)

	routesDef := gin.Default()

	// routes
	routesDef = routes.UserRouter(userHandler, authService, userService)

	routesDef.Run()
}
