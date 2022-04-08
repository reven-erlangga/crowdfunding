package formatter

import "crowdfunding-server/models"

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token"`
}

func FormatUser(user models.User, token string) UserFormatter {
	return UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Password:   user.PasswordHash,
		Token:      token,
	}
}
