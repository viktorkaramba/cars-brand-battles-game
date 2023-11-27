package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
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

func (r *ScorePostgres) Update(id int) error {

	query := fmt.Sprintf("UPDATE %s s SET playerScore=playerScore+1 WHERE s.id=$1", scoresTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ScorePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", scoresTable)
	_, err := r.db.Exec(query, id)
	return err
}
