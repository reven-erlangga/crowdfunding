package handlers

import (
	"crowdfunding-server/formatter"
	"crowdfunding-server/helpers"
	"crowdfunding-server/requests"
	"crowdfunding-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

// Implementasi todo handler

func (h *userHandler) Create(ctx *gin.Context) {
	var request requests.UserRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		errorMessage := gin.H{
			"errors": helpers.ValidationError(err),
		}

		response := helpers.ApiResponse("Register account failed!", http.StatusUnprocessableEntity, "error", errorMessage)

		ctx.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	newUser, err := h.userService.Create(request)

	if err != nil {
		response := helpers.ApiResponse("Register account failed!", http.StatusBadRequest, "error", nil)

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := formatter.FormatUser(newUser, "")

	response := helpers.ApiResponse("Account created successfully", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
