package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"strings"
)

type ScorePostgres struct {
	db *sqlx.DB
}

func NewScorePostgres(db *sqlx.DB) *ScorePostgres {
	return &ScorePostgres{db: db}
}

func (r *ScorePostgres) GetAll() ([]carsBrandsBattle.Score, error) {
	var scores []carsBrandsBattle.Score
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC ", scoresTable)
	err := r.db.Select(&scores, query)
	return scores, err
}

func (r *ScorePostgres) GetById(id int) (*carsBrandsBattle.Score, error) {
	var score carsBrandsBattle.Score
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 ORDER BY id ASC ", scoresTable)
	err := r.db.Get(&score, query, id)
	if score == (carsBrandsBattle.Score{}) {
		return &carsBrandsBattle.Score{}, nil
	}
	return &score, err
}

func (r *ScorePostgres) Update(id int, score carsBrandsBattle.UpdateScoreInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if score.PlayerScore != nil {
		setValues = append(setValues, fmt.Sprintf("playerScore=$%d", argId))
		args = append(args, *score.PlayerScore)
		argId++
	}

	if score.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("userId=$%d", argId))
		args = append(args, *score.UserId)
		argId++
	}

	if score.BattleId != nil {
		setValues = append(setValues, fmt.Sprintf("battleId=$%d", argId))
		args = append(args, *score.BattleId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s s SET %s WHERE s.id=$%d", scoresTable, setQuery, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ScorePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", scoresTable)
	_, err := r.db.Exec(query, id)
	return err
}
