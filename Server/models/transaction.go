package models

import (
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       User
	Campaign   Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
