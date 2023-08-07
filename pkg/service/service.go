package service

import (
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user carsBrandRandomGenerator.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Brand interface {
}

type Battle interface {
}

type Score interface {
}

type Service struct {
	Authorization
	Battle
	Score
	Brand
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
