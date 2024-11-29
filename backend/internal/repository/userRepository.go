package repository

import (
	"fmt"

	"backend/internal/models"
)

//go:generate mockgen -destination=mocks/mockUserRepository.go -package=mocks backend/internal/repository IUserRepository
type IUserRepository interface {
	SaveUser(user models.User) (models.User, error)
	GetUser(username string) (models.User, error)
}

type UserRepository struct {
}

func ProvideUserRepository() IUserRepository {
	return &UserRepository{}
}

var users map[string]models.User

func (ur UserRepository) SaveUser(user models.User) (models.User, error) {
	if users == nil {
		users = make(map[string]models.User)
	}

	if _, isFound := users[user.Username]; isFound {
		return models.User{}, fmt.Errorf("user already exists")
	}

	users[user.Username] = user

	return user, nil
}

func (ur UserRepository) GetUser(username string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}
