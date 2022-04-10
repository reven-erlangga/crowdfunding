package handlers

import (
	"crowdfunding-server/formatter"
	"crowdfunding-server/helpers"
	"crowdfunding-server/models"
	"crowdfunding-server/requests"
	"crowdfunding-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(userService services.TransactionService) *transactionHandler {
	return &transactionHandler{userService}
}

func (h *transactionHandler) GetCampaignTransactions(ctx *gin.Context) {
	var request requests.GetTransactionsCampaignRequest

	err := ctx.ShouldBindUri(&request)

	if err != nil {
		response := helpers.ApiResponse("Failed to get campaign transactions!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionByCampaignID(request)

	if err != nil {
		response := helpers.ApiResponse("Failed to get campaign transactions!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("List of campaign transactions", http.StatusOK, "success", formatter.FormatCampaignTransactions(transactions))
	ctx.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetUserTransactions(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionByUserID(userID)

	if err != nil {
		response := helpers.ApiResponse("Failed to get user transactions!", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("List of user transactions", http.StatusOK, "success", formatter.FormatUserTransactions(transactions))
	ctx.JSON(http.StatusOK, response)

}
