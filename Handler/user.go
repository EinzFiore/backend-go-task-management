package handler

import (
	users "crowdfunding/Users"
	"errors"
	"net/http"

	helpers "crowdfunding/Helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterInput

	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("", http.StatusUnprocessableEntity, nil)

	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response.Meta.Message = "Failed register account"
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser := users.UserFormat(user, "newtokenuser")

	response = helpers.ApiResponse("Account has been registered", http.StatusOK, newUser)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input users.LoginInput
	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("Opps, something error", http.StatusBadRequest, nil)

	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)
		response.Meta.Code = http.StatusUnprocessableEntity
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogged, err := h.userService.Login(input)
	if err != nil {
		response.Meta.Errors = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userForm := users.UserFormat(userLogged, "tokenlogin")
	response = helpers.ApiResponse("Login succes", http.StatusOK, userForm)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	response := helpers.ApiResponse("Opps, something error", http.StatusBadRequest, nil)
	if err != nil {
		response.Meta.Errors = helpers.ValidatorError(err)
		response.Meta.Code = http.StatusUnprocessableEntity
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resEmail, err := h.userService.IsEmailAvailable(input)
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
