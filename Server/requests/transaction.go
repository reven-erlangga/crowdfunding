package requests

import "crowdfunding-server/models"

type GetTransactionsCampaignRequest struct {
	ID   int         `uri:"id" binding:"required"`
	User models.User `json:"user"`
}
