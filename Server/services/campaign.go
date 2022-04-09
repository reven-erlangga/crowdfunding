package services

import (
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
	"crowdfunding-server/requests"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]models.Campaign, error)
	GetCampaignByID(request requests.GetCampaignDetailRequest) (models.Campaign, error)
	CreateCampaign(request requests.CreateCampaignRequest) (models.Campaign, error)
	UpdateCampaign(requestID requests.GetCampaignDetailRequest, requestData requests.CreateCampaignRequest) (models.Campaign, error)
}

type campaignService struct {
	repository repositories.CampaignRepository
}

func NewCampaignService(repository repositories.CampaignRepository) *campaignService {
	return &campaignService{repository}
}

func (s *campaignService) GetCampaigns(userID int) ([]models.Campaign, error) {
	if userID != 0 {
		campaign, err := s.repository.FindByUserID(userID)

		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *campaignService) GetCampaignByID(request requests.GetCampaignDetailRequest) (models.Campaign, error) {
	campaign, err := s.repository.FindByID(request.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *campaignService) CreateCampaign(request requests.CreateCampaignRequest) (models.Campaign, error) {
	slugCandidate := fmt.Sprintf("%s %d", request.Name, request.User.ID)
	slugCandidate = slug.Make(slugCandidate)

	campaign := models.Campaign{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		Description:      request.Description,
		GoalAmount:       request.GoalAmount,
		Perks:            request.Perks,
		UserID:           request.User.ID,
		Slug:             slugCandidate,
	}

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, nil
	}

	return newCampaign, nil
}

func (s *campaignService) UpdateCampaign(requestID requests.GetCampaignDetailRequest, requestData requests.CreateCampaignRequest) (models.Campaign, error) {
	campaign, err := s.repository.FindByID(requestID.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.UserID != requestData.User.ID {
		return campaign, errors.New("You are not allowed to update this campaign")
	}

	campaign.Name = requestData.Name
	campaign.ShortDescription = requestData.ShortDescription
	campaign.Description = requestData.Description
	campaign.GoalAmount = requestData.GoalAmount
	campaign.Perks = requestData.Perks

	updatedCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}
