package handlers

import (
	"crowdfunding-server/formatter"
	"crowdfunding-server/helpers"
	"crowdfunding-server/requests"
	"crowdfunding-server/services"
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
