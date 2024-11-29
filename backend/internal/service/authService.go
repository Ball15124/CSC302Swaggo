package service

import (
	"fmt"

	"backend/internal/models"
	"backend/internal/repository"
)

type IAuthService interface {
	Register(username string, password string, email string) (models.User, error)
	Login(username string, password string) (models.User, error)
	GetUser(username string) (models.User, error)
}

type AuthService struct {
	hashService IHashService
	userRepo    repository.IUserRepository
}

func ProvideAuthService(hs IHashService, ur repository.IUserRepository) IAuthService {
	return &AuthService{
		hashService: hs,
		userRepo:    ur,
	}
}

func (a AuthService) Register(username string, password string, email string) (models.User, error) {
	hashedPassword, err := a.hashService.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}
	return a.userRepo.SaveUser(user)
}

func (a AuthService) Login(username string, password string) (models.User, error) {
	user, err := a.userRepo.GetUser(username)
	if err != nil {
		return models.User{}, err
	}
	if !a.hashService.ComparePassword(user.Password, password) {
		return models.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}

func (a AuthService) GetUser(username string) (models.User, error) {
	user, err := a.userRepo.GetUser(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
