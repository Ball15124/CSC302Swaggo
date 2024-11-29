package service

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -destination=./mock/mockIHashService.go -package=mock backend/internal/service IHashService
type IHashService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
	Base64Encode(input string) (string, error)
}

type HashService struct {
	salt int
}

func ProvideHashService() IHashService {
	return &HashService{
		salt: 10,
	}
}

func (hs *HashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hs.salt)
	return string(bytes), err
}

func (hs *HashService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (hs *HashService) Base64Encode(input string) (string, error) {
	data := base64.StdEncoding.EncodeToString([]byte(input))
	return data, nil
}
