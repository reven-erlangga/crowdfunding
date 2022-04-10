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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
