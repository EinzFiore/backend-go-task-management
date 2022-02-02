package routes

import (
	handler "crowdfunding/Handler/Users"
	middleware "crowdfunding/Middleware"
	auth "crowdfunding/Services/Auth"
	users "crowdfunding/Services/Users"

	"github.com/gin-gonic/gin"
)

func UserRouter(userHandler handler.UserHandler, authService auth.Service, userService users.Service) *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/checkEmail", userHandler.CheckEmailAvailability)
	api.POST("/avatar", middleware.Middleware(authService, userService), userHandler.UploadAvatar)

	return router
}
