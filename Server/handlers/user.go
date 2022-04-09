package handlers

import (
	"crowdfunding-server/formatter"
	"crowdfunding-server/helpers"
	"crowdfunding-server/models"
	"crowdfunding-server/requests"
	"crowdfunding-server/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
	authService services.AuthService
}

func NewUserHandler(userService services.UserService, authService services.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

// Implementasi todo handler

func (h *userHandler) Create(ctx *gin.Context) {
	var request requests.RegisterUserRequest

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
		response := helpers.ApiResponse("Register account failed!", http.StatusBadRequest, "error", err.Error())

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helpers.ApiResponse("Register account failed!", http.StatusBadRequest, "error", err.Error())

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := formatter.FormatUser(newUser, token)

	response := helpers.ApiResponse("Account created successfully", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var request requests.LoginUserRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		errorMessage := gin.H{
			"errors": helpers.ValidationError(err),
		}

		response := helpers.ApiResponse("Login failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(request)

	if err != nil {
		response := helpers.ApiResponse("Login failed!", http.StatusUnprocessableEntity, "error", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		response := helpers.ApiResponse("Login failed!", http.StatusBadRequest, "error", err.Error())

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	formatUser := formatter.FormatUser(loggedInUser, token)

	response := helpers.ApiResponse("Login successfully", http.StatusOK, "success", formatUser)

	ctx.JSON(http.StatusOK, response)
}

// Check ketersediaan email
func (h *userHandler) CheckEmailAvailability(ctx *gin.Context) {
	var request requests.CheckEmailRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helpers.ApiResponse("Email checking failed!", http.StatusUnprocessableEntity, "error", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(request)

	if err != nil {
		response := helpers.ApiResponse("Email checking failed!", http.StatusUnprocessableEntity, "error", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helpers.ApiResponse(metaMessage, http.StatusOK, "success", data)
	ctx.JSON(http.StatusOK, response)
}

// Upload avatar
func (h *userHandler) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse("Upload avatar failed!", http.StatusUnprocessableEntity, "error", data)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/avatars/%d-%s", userId, file.Filename)

	err = ctx.SaveUploadedFile(file, "assets/"+path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse("Upload avatar failed!", http.StatusUnprocessableEntity, "error", data)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse("Upload avatar failed!", http.StatusUnprocessableEntity, "error", data)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helpers.ApiResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)
}
