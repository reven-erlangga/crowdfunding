package models

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentMounth    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
