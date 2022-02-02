package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	CheckEmailAvailability(c *gin.Context)
	UploadAvatar(c *gin.Context)
}
