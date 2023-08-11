package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type UserStatisticsPostgres struct {
	db *sqlx.DB
}

func NewUserStatisticsPostgres(db *sqlx.DB) *UserStatisticsPostgres {
	return &UserStatisticsPostgres{db: db}
}

func (r *UserStatisticsPostgres) GetGeneralStatisticsByScore() ([]carsBrandsBattle.UserStatistics, error) {
	var userStatistics []carsBrandsBattle.UserStatistics
	query := fmt.Sprintf(
		"SELECT users.id, users.username, SUM(scores.playerScore) FROM %s JOIN %s ON users.id=scores.userId "+
			"GROUP BY users.id order by SUM(scores.playerScore) DESC", usersTable, scoresTable)
	err := r.db.Select(&userStatistics, query)
	return userStatistics, err
}
