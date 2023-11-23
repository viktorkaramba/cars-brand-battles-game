package repository

import (
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type Authorization interface {
	CreateUser(user carsBrandsBattle.User) (int, error)
	GetUser(username, password string) (*carsBrandsBattle.User, error)
	GetUserByUsername(username string) (*carsBrandsBattle.User, error)
}

type Brand interface {
	Create(brand carsBrandsBattle.Brand) (int, error)
	GetAll() ([]carsBrandsBattle.Brand, error)
	GetById(id int) (*carsBrandsBattle.Brand, error)
	Update(id int, brand carsBrandsBattle.UpdateBrandInput) error
	Delete(id int) error
}

type Battle interface {
	Create(battle carsBrandsBattle.Battle) (int, error)
	GetAll() ([]carsBrandsBattle.Battle, error)
	GetById(id int) (*carsBrandsBattle.Battle, error)
	Update(id int, battle carsBrandsBattle.UpdateBattleInput) error
	Delete(id int) error
}

type Score interface {
	GetAll() ([]carsBrandsBattle.Score, error)
	GetById(id int) (*carsBrandsBattle.Score, error)
	Update(id int, score carsBrandsBattle.UpdateScoreInput) error
	Delete(id int) error
}

type Tokens interface {
	Create(token carsBrandsBattle.Token) (int, error)
	GetByToken(token string) (*carsBrandsBattle.Token, error)
	Update(token string, updatedToken carsBrandsBattle.UpdateTokenInput) error
}

type UserStatistics interface {
	GetGeneralStatisticsByScore() ([]carsBrandsBattle.UserStatistics, error)
}

type UserInterfaceData interface {
	GetAll(isFinished bool) ([]carsBrandsBattle.UserInterfaceData, error)
	GetById(battleId int, isFinished bool) (*carsBrandsBattle.UserInterfaceData, error)
}

type Repository struct {
	Authorization
	Brand
	Battle
	Score
	Tokens
	UserStatistics
	UserInterfaceData
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:     NewAuthPostgres(db),
		Battle:            NewBattlePostgres(db),
		Brand:             NewBrandPostgres(db),
		Score:             NewScorePostgres(db),
		Tokens:            NewTokenPostgres(db),
		UserStatistics:    NewUserStatisticsPostgres(db),
		UserInterfaceData: NewUserInterfaceDataPostgres(db),
	}
}
