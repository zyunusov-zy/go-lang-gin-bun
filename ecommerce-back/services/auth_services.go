package services

import (
	"ecommerce-back/models"
	"ecommerce-back/repositories"
	"ecommerce-back/utils"
	"errors"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(user *models.User) error {
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd
	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (*models.User, error){
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}