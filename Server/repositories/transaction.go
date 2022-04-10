package repositories

import (
	"crowdfunding-server/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetCampaignByID(campaignID int) ([]models.Transaction, error)
	GetByUserID(userID int) ([]models.Transaction, error)
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

func (r *transactionRepository) GetByUserID(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id DESC").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
