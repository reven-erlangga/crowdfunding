package repositories

import (
	"errors"

	"gorm.io/gorm"

	"crowdfunding-server/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(ID int) (models.User, error)
	Create(todo models.User) (models.User, error)
	Update(todo models.User) (models.User, error)
	Delete(todo models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Implementasi user repository

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindByID(ID int) (models.User, error) {
	var user models.User

	result := r.db.Find(&user, ID)

	err := result.Error

	// Apabila data tidak ada / tidak ditemukan maka return error
	if result.RowsAffected == 0 && err == nil {
		err = errors.New("todo not found")
	}

	return user, err
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) Update(todo models.User) (models.User, error) {
	err := r.db.Save(&todo).Error

	return todo, err
}

func (r *userRepository) Delete(todo models.User) (models.User, error) {
	err := r.db.Delete(&todo).Error

	return todo, err
}
