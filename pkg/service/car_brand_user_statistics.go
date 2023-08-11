package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type UserStatisticService struct {
	repo repository.UserStatistics
}

func NewUserStatisticService(repo repository.UserStatistics) *UserStatisticService {
	return &UserStatisticService{
		repo: repo,
	}
}

func (s *UserStatisticService) GetGeneralStatisticsByScore() ([]carsBrandsBattle.UserStatistics, error) {
	return s.repo.GetGeneralStatisticsByScore()
}
