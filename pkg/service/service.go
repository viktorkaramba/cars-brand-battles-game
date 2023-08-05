package service

import "github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"

type Brand interface {
}

type Service struct {
	Brand
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
