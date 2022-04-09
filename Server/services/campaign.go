package services

import (
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]models.Campaign, error)
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
