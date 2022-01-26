package handler

import (
	users "crowdfunding/Users"
	"net/http"

	helpers "crowdfunding/Helpers"

	"github.com/gin-gonic/gin"
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
