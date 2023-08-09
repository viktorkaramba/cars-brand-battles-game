package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type ScorePostgres struct {
	db *sqlx.DB
}

func NewScorePostgres(db *sqlx.DB) *ScorePostgres {
	return &ScorePostgres{db: db}
}

func (r *ScorePostgres) GetAll() ([]carsBrandRandomGenerator.Score, error) {
	var scores []carsBrandRandomGenerator.Score
	query := fmt.Sprintf("SELECT * FROM %s", scoresTable)
	err := r.db.Select(&scores, query)
	return scores, err
}

func (r *ScorePostgres) GetById(id int) (carsBrandRandomGenerator.Score, error) {
	var score carsBrandRandomGenerator.Score
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", scoresTable)
	err := r.db.Select(&score, query, id)
	return score, err
}

func (r *ScorePostgres) Update(id int, score carsBrandRandomGenerator.UpdateScoreInput) error {
	query := fmt.Sprintf("UPDATE %s SET userId=$1, battleId=$2, playerScore=$3 WHERE id=$4", brandsTable)
	_, err := r.db.Exec(query, score.UserId, score.BattleId, score.PlayerScore, id)
	return err
}

func (r *ScorePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", scoresTable)
	_, err := r.db.Exec(query, id)
	return err
}
