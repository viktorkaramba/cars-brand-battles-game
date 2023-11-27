package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

type BattleService struct {
	repo repository.Battle
}

func NewBattleService(repo repository.Battle) *BattleService {
	return &BattleService{
		repo: repo,
	}
}

func (s *BattleService) Create(battle carsBrandsBattle.Battle) (int, error) {
	return s.repo.Create(battle)
}

func (s *BattleService) GetAll() ([]carsBrandsBattle.Battle, error) {
	return s.repo.GetAll()
}

func (s *BattleService) GetById(id int) (*carsBrandsBattle.Battle, error) {
	return s.repo.GetById(id)
}

func (s *BattleService) Update(id int) error {
	return s.repo.Update(id)
}

func (s *BattleService) Delete(id int) error {
	return s.repo.Delete(id)
}
