package service

import (
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
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

func (s *BattleService) Create(battle carsBrandRandomGenerator.Battle) (int, error) {
	return s.repo.Create(battle)
}

func (s *BattleService) GetAll() ([]carsBrandRandomGenerator.Battle, error) {
	return s.repo.GetAll()
}

func (s *BattleService) GetById(id int) (carsBrandRandomGenerator.Battle, error) {
	return s.repo.GetById(id)
}

func (s *BattleService) Update(id int, battle carsBrandRandomGenerator.UpdateBattleInput) error {
	if err := battle.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, battle)
}

func (s *BattleService) Delete(id int) error {
	return s.repo.Delete(id)
}
