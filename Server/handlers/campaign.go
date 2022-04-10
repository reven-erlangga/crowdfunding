package handlers

import (
	"crowdfunding-server/formatter"
	"crowdfunding-server/helpers"
	"crowdfunding-server/models"
	"crowdfunding-server/requests"
	"crowdfunding-server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service services.CampaignService
}

func NewCampaignHandler(userService services.CampaignService) *campaignHandler {
	return &campaignHandler{userService}
}

// Implementasi handler

func (h *campaignHandler) GetCampaigns(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helpers.ApiResponse("Error to get campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("List of campaign", http.StatusOK, "success", formatter.FormatCampaigns(campaigns))
	ctx.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(ctx *gin.Context) {
	var request requests.GetCampaignDetailRequest

	err := ctx.ShouldBindUri(&request)
	fmt.Println(request.ID)
	if err != nil {
		response := helpers.ApiResponse("Failed to get campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.service.GetCampaignByID(request)

	if err != nil {
		response := helpers.ApiResponse("Failed to get detail campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("Detail campaign", http.StatusOK, "success", formatter.FormatCampaignDetail(campaign))
	ctx.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(ctx *gin.Context) {
	var request requests.CreateCampaignRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.ValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helpers.ApiResponse("Failed to create campaign!", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)

	request.User = currentUser

	newCampaign, err := h.service.CreateCampaign(request)

	if err != nil {
		response := helpers.ApiResponse("Failed to create campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("Success to create campaign!", http.StatusCreated, "success", formatter.FormatCampaignDetail(newCampaign))
	ctx.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdateCampaign(ctx *gin.Context) {
	var requestID requests.GetCampaignDetailRequest

	err := ctx.ShouldBindUri(&requestID)

	if err != nil {
		response := helpers.ApiResponse("Failed to update campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var requestData requests.CreateCampaignRequest

	err = ctx.ShouldBindJSON(&requestData)

	if err != nil {
		errors := helpers.ValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helpers.ApiResponse("Failed to update campaign!", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)
	requestData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(requestID, requestData)

	if err != nil {
		response := helpers.ApiResponse("Failed to update campaign!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return

	}

	response := helpers.ApiResponse("Success to update campaign!", http.StatusOK, "success", formatter.FormatCampaign(updateCampaign))
	ctx.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(ctx *gin.Context) {
	var request requests.CreateCampaignImageRequest

	err := ctx.ShouldBind(&request)

	if err != nil {
		errors := helpers.ValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helpers.ApiResponse("Failed to upload image!", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := ctx.MustGet("currentUser").(models.User)

	request.User = currentUser

	file, err := ctx.FormFile("file")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse("Failed to upload campaign image!", http.StatusBadRequest, "error", data)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/campaigns/%d-%s", request.User.ID, file.Filename)

	err = ctx.SaveUploadedFile(file, "assets/"+path)

	if err != nil {
		data := gin.H{
			"errors": err.Error(),
		}
		response := helpers.ApiResponse("Cant save image!", http.StatusBadRequest, "error", data)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(request, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse("Upload campaign image failed!", http.StatusBadRequest, "error", data)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helpers.ApiResponse("Upload campaign success!", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)
}
