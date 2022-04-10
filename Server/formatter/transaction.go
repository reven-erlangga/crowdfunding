package formatter

import (
	"crowdfunding-server/models"
	"time"
)

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction models.Transaction) TransactionFormatter {
	return TransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

func FormatCampaignTransactions(transactions []models.Transaction) []TransactionFormatter {
	if len(transactions) == 0 {
		return []TransactionFormatter{}
	}

	var transactionsFormatter []TransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int                          `json:"id"`
	Amount    int                          `json:"amount"`
	Status    string                       `json:"status"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt time.Time                    `json:"updated_at"`
	Campaign  CampaignTransactionFormatter `json:"campaign"`
}

type CampaignTransactionFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction models.Transaction) UserTransactionFormatter {
	imageURL := ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		imageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	campaignTransactionFormat := CampaignTransactionFormatter{
		Name:     transaction.Campaign.Name,
		ImageURL: imageURL,
	}

	return UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  campaignTransactionFormat,
	}

}

func FormatUserTransactions(transactions []models.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, formatter)
	}

	return userTransactionsFormatter
}
