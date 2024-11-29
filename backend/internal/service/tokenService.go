package service

import (
	"fmt"

	uuid2 "github.com/google/uuid"
)

type ITokenService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

// this is wrong, but it's just for the sake of the example
// don't store tokens in memory, use a database or a cache and set an expiration time
// or use JWT
type TokenService struct {
	secretKey string
}

func ProvideTokenService() ITokenService {
	return TokenService{}
}

var tokens = make(map[string]string)

func (t TokenService) GenerateToken(username string) (string, error) {
	uuid := uuid2.NewString()
	tokens[uuid] = username

	return uuid, nil
}

func (t TokenService) ValidateToken(token string) (string, error) {
	username, isFound := tokens[token]
	if !isFound {
		return "", fmt.Errorf("invalid token")
	}

	return username, nil
}
