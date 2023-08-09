package repository

import (
	"github.com/jmoiron/sqlx"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type Authorization interface {
	CreateUser(user carsBrandRandomGenerator.User) (int, error)
	GetUser(username, password string) (carsBrandRandomGenerator.User, error)
}

type Brand interface {
	Create(brand carsBrandRandomGenerator.Brand) (int, error)
	GetAll() ([]carsBrandRandomGenerator.Brand, error)
	GetById(id int) (carsBrandRandomGenerator.Brand, error)
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

type Repository struct {
	Authorization
	Brand
	Battle
	Score
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Battle:        NewBattlePostgres(db),
		Brand:         NewBrandPostgres(db),
		Score:         NewScorePostgres(db),
	}
}
