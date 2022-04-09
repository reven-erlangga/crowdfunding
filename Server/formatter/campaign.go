package formatter

import (
	"crowdfunding-server/models"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign models.Campaign) CampaignFormatter {
	imageUrl := ""

	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	return CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageUrl:         imageUrl,
	}
}

func FormatCampaigns(campaigns []models.Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)

		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign models.Campaign) CampaignDetailFormatter {
	imageUrl := ""

	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter := CampaignImageFormatter{
			ImageURL:  image.FileName,
			IsPrimary: isPrimary,
		}

		images = append(images, campaignImageFormatter)
	}

	return CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		ImageUrl:         imageUrl,
		Perks:            perks,
		User:             campaignUserFormatter,
		Images:           images,
	}
}
