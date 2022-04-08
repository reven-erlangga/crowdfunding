package services

import (
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
	"crowdfunding-server/requests"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(userRequest requests.UserRequest) (models.User, error)
}

type userServices struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userServices {
	return &userServices{userRepository}
}

// Implementasi todo handler

func (s *userServices) Create(request requests.UserRequest) (models.User, error) {
	user := models.User{}

	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.userRepository.Create(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
