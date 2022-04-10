package services

import (
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
	"crowdfunding-server/requests"
	"errors"
)

type TransactionService interface {
	GetTransactionByCampaignID(request requests.GetTransactionsCampaignRequest) ([]models.Transaction, error)
	GetTransactionByUserID(userID int) ([]models.Transaction, error)
}

type transactionService struct {
	repository         repositories.TransactionRepository
	campaignRepository repositories.CampaignRepository
}

func NewTransactionService(repository repositories.TransactionRepository, campaignRepository repositories.CampaignRepository) *transactionService {
	return &transactionService{repository, campaignRepository}
}

func (s *transactionService) GetTransactionByCampaignID(request requests.GetTransactionsCampaignRequest) ([]models.Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(request.ID)

	if err != nil {
		return []models.Transaction{}, err
	}

	if campaign.User.ID != request.User.ID {
		return []models.Transaction{}, errors.New("you are not authorized to view this campaign")
	}

	transactions, err := s.repository.GetCampaignByID(request.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionByUserID(userID int) ([]models.Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
