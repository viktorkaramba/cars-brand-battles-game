package service

import (
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
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

func (s *ScoreService) GetAll() ([]carsBrandRandomGenerator.Score, error) {
	return s.repo.GetAll()
}

func (s *ScoreService) GetById(id int) (carsBrandRandomGenerator.Score, error) {
	return s.repo.GetById(id)
}

func (s *ScoreService) Update(id int, score carsBrandRandomGenerator.UpdateScoreInput) error {
	if err := score.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, score)
}

func (s *ScoreService) Delete(id int) error {
	return s.repo.Delete(id)
}
