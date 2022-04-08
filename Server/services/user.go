package services

import (
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
	"crowdfunding-server/requests"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(userRequest requests.RegisterUserRequest) (models.User, error)
	Login(userRequest requests.LoginUserRequest) (models.User, error)
}

type userServices struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userServices {
	return &userServices{userRepository}
}

// Implementasi todo handler

func (s *userServices) Create(request requests.RegisterUserRequest) (models.User, error) {
	user := models.User{}

	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	userExist, _ := s.userRepository.FindByEmail(request.Email)

	if userExist.ID != 0 {
		return user, errors.New("user already exist")
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.userRepository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userServices) Login(request requests.LoginUserRequest) (models.User, error) {
	email := request.Email
	password := request.Password

	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
