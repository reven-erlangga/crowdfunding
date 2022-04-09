package requests

import "crowdfunding-server/models"

type GetCampaignDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignRequest struct {
	Name             string      `json:"name" binding:"required"`
	ShortDescription string      `json:"short_description" binding:"required"`
	Description      string      `json:"description" binding:"required"`
	GoalAmount       int         `json:"goal_amount" binding:"required"`
	Perks            string      `json:"perks"`
	User             models.User `json:"user"`
}
