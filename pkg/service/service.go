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
	Create(brand carsBrandRandomGenerator.Brand) (int, error)
	GetAll() ([]carsBrandRandomGenerator.Brand, error)
	GetById(id int) (carsBrandRandomGenerator.Brand, error)
	GetOneByRandom() (carsBrandRandomGenerator.Brand, error)
	Update(id int, brand carsBrandRandomGenerator.UpdateBrandInput) error
	Delete(id int) error
}

type Battle interface {
	Create(battle carsBrandRandomGenerator.Battle) (int, error)
	GetAll() ([]carsBrandRandomGenerator.Battle, error)
	GetById(id int) (carsBrandRandomGenerator.Battle, error)
	Update(id int, battle carsBrandRandomGenerator.UpdateBattleInput) error
	Delete(id int) error
}

type Score interface {
	GetAll() ([]carsBrandRandomGenerator.Score, error)
	GetById(id int) (carsBrandRandomGenerator.Score, error)
	Update(id int, score carsBrandRandomGenerator.UpdateScoreInput) error
	Delete(id int) error
}

type Service struct {
	Authorization
	Battle
	Score
	Brand
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Battle:        NewBattleService(repos.Battle),
		Score:         NewScoreService(repos.Score),
		Brand:         NewBrandService(repos.Brand),
	}
}
