package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type TokenService struct {
	repo repository.Tokens
}

func NewTokenService(repo repository.Tokens) *TokenService {
	return &TokenService{
		repo: repo,
	}
}

func (s *TokenService) Create(token carsBrandsBattle.Token) (int, error) {
	return s.repo.Create(token)
}

func (s *TokenService) GetByToken(token string) (*carsBrandsBattle.Token, error) {
	return s.repo.GetByToken(token)
}

func (s *TokenService) Update(token string, updatedToken carsBrandsBattle.UpdateTokenInput) error {
	if err := updatedToken.Validate(); err != nil {
		return err
	}
	return s.repo.Update(token, updatedToken)
}
