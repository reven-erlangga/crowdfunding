package repositories

import (
	"crowdfunding-server/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetCampaignByID(campaignID int) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetCampaignByID(campaignID int) ([]models.Transaction, error) {
	transactions := []models.Transaction{}

	if err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id DESC").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
