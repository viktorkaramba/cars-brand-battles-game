package repository

import (
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type Authorization interface {
	CreateUser(user carsBrandsBattle.User) (int, error)
	GetUser(username, password string) (carsBrandsBattle.User, error)
}

type Brand interface {
	Create(brand carsBrandsBattle.Brand) (int, error)
	GetAll() ([]carsBrandsBattle.Brand, error)
	GetById(id int) (carsBrandsBattle.Brand, error)
	Update(id int, brand carsBrandsBattle.UpdateBrandInput) error
	Delete(id int) error
}

type Battle interface {
	Create(battle carsBrandsBattle.Battle) (int, error)
	GetAll() ([]carsBrandsBattle.Battle, error)
	GetById(id int) (carsBrandsBattle.Battle, error)
	Update(id int, battle carsBrandsBattle.UpdateBattleInput) error
	Delete(id int) error
}

type Score interface {
	GetAll() ([]carsBrandsBattle.Score, error)
	GetById(id int) (carsBrandsBattle.Score, error)
	Update(id int, score carsBrandsBattle.UpdateScoreInput) error
	Delete(id int) error
}

type User interface {
	GetAll() ([]carsBrandsBattle.User, error)
	GetById(id int) (carsBrandsBattle.Score, error)
	Update(id int, score carsBrandsBattle.UpdateScoreInput) error
	Delete(id int) error
}

type UserStatistics interface {
	GetGeneralStatisticsByScore() ([]carsBrandsBattle.UserStatistics, error)
}

type UserInterfaceData interface {
	GetAll() ([]carsBrandsBattle.UserInterfaceData, error)
}

type Repository struct {
	Authorization
	Brand
	Battle
	Score
	UserStatistics
	UserInterfaceData
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:     NewAuthPostgres(db),
		Battle:            NewBattlePostgres(db),
		Brand:             NewBrandPostgres(db),
		Score:             NewScorePostgres(db),
		UserStatistics:    NewUserStatisticsPostgres(db),
		UserInterfaceData: NewUserInterfaceDataPostgres(db),
	}
}
