package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user carsBrandsBattle.User) (int, error)
	GetUserByUsername(username string) (*carsBrandsBattle.User, error)
	GenerateToken(username, password string) (*carsBrandsBattle.User, string, error)
	RefreshToken(userId int) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Brand interface {
	Create(brand carsBrandsBattle.Brand) (int, error)
	GetAll() ([]carsBrandsBattle.Brand, error)
	GetById(id int) (*carsBrandsBattle.Brand, error)
	GetOneByRandom() (carsBrandsBattle.Brand, error)
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

type Service struct {
	Authorization
	Battle
	Score
	Brand
	Tokens
	UserStatistics
	UserInterfaceData
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:     NewAuthService(repos.Authorization),
		Battle:            NewBattleService(repos.Battle),
		Score:             NewScoreService(repos.Score),
		Brand:             NewBrandService(repos.Brand),
		Tokens:            NewTokenService(repos.Tokens),
		UserStatistics:    NewUserStatisticService(repos.UserStatistics),
		UserInterfaceData: NewUserInterfaceDataService(repos.UserInterfaceData),
	}
}
