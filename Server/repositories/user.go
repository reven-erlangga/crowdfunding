package repositories

import (
	"gorm.io/gorm"

	"crowdfunding-server/models"
)

type UserRepository interface {
	Save(todo models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Implementasi user repository

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
