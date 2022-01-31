package handler

import (
	auth "crowdfunding/Auth"
	users "crowdfunding/Users"
	"errors"
	"fmt"
	"net/http"

	helpers "crowdfunding/Helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	userService users.Service
	authService auth.Service
}

func NewUserHandler(userService users.Service, authServuce auth.Service) *userHandler {
	return &userHandler{userService, authServuce}
}

func (service *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterInput

	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("", http.StatusUnprocessableEntity, nil)

	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := service.userService.RegisterUser(input)
	if err != nil {
		response.Meta.Message = "Failed register account"
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := service.authService.GenerateToken(user.Id)
	if err != nil {
		response.Meta.Message = "Failed while generate token"
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	newUser := users.UserFormat(user, token)

	response = helpers.ApiResponse("Account has been registered", http.StatusOK, newUser)

	c.JSON(http.StatusOK, response)
}

func (service *userHandler) Login(c *gin.Context) {
	var input users.LoginInput
	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("Opps, something error", http.StatusBadRequest, nil)

	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)
		response.Meta.Code = http.StatusUnprocessableEntity
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogged, err := service.userService.Login(input)
	if err != nil {
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := service.authService.GenerateToken(userLogged.Id)
	if err != nil {
		response.Meta.Message = "Failed while generate token"
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userForm := users.UserFormat(userLogged, token)
	response = helpers.ApiResponse("Login succes", http.StatusOK, userForm)
	c.JSON(http.StatusOK, response)
}

func (service *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("Opps, something error", http.StatusBadRequest, nil)
	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)
		response.Meta.Code = http.StatusUnprocessableEntity
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resEmail, err := service.userService.IsEmailAvailable(input)
	if err != nil {
		response.Meta.Errors = err.Error()
		if errors.Is(err, gorm.ErrRecordNotFound) == false {
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	metaMessage := "email is available"
	if resEmail {
		metaMessage = "Email has been registered"
	}

	response = helpers.ApiResponse(metaMessage, http.StatusOK, resEmail)

	c.JSON(http.StatusOK, response)
}

func (service *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	response := helpers.ApiResponse("Opps, something error", http.StatusBadRequest, nil)

	if err != nil {
		response.Meta.Errors = err.Error()
		response.Data = false
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// hardcode user id, next will get from JWT
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		response.Meta.Message = "Failed to upload avatar"
		response.Meta.Errors = err.Error()
		response.Data = false

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = service.userService.SaveAvatar(userID, path)
	if err != nil {
		response.Meta.Message = "Failed to upload avatar"
		response.Meta.Errors = err.Error()
		response.Data = false

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = helpers.ApiResponse("Success upload avatar", http.StatusOK, true)
	c.JSON(http.StatusOK, response)
}
