package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type ScoreService struct {
	repo repository.Score
}

func NewScoreService(repo repository.Score) *ScoreService {
	return &ScoreService{
		repo: repo,
	}
}

func (s *ScoreService) GetAll() ([]carsBrandsBattle.Score, error) {
	return s.repo.GetAll()
}

func (s *ScoreService) GetById(id int) (*carsBrandsBattle.Score, error) {
	return s.repo.GetById(id)
}

func (s *ScoreService) Update(id int) error {
	return s.repo.Update(id)
}

func (s *ScoreService) Delete(id int) error {
	return s.repo.Delete(id)
}
