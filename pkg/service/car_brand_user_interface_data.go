package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type UserInterfaceDataService struct {
	repo repository.UserInterfaceData
}

func NewUserInterfaceDataService(repo repository.UserInterfaceData) *UserInterfaceDataService {
	return &UserInterfaceDataService{
		repo: repo,
	}
}

func (s *UserInterfaceDataService) GetAll(isFinished bool) ([]carsBrandsBattle.UserInterfaceData, error) {
	return s.repo.GetAll(isFinished)
}

func (s *UserInterfaceDataService) GetById(battleId int, isFinished bool) (*carsBrandsBattle.UserInterfaceData, error) {
	return s.repo.GetById(battleId, isFinished)
}
