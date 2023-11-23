package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"strings"
)

type TokenPostgres struct {
	db *sqlx.DB
}

func NewTokenPostgres(db *sqlx.DB) *TokenPostgres {
	return &TokenPostgres{db: db}
}

func (r *TokenPostgres) Create(token carsBrandsBattle.Token) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (tokenValue, revoked, userId) values ($1, $2, $3) RETURNING id",
		tokensTable)
	row := r.db.QueryRow(query, token.TokenValue, token.Revoked, token.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TokenPostgres) GetByToken(token string) (*carsBrandsBattle.Token, error) {
	var newToken carsBrandsBattle.Token
	query := fmt.Sprintf("SELECT * FROM %s WHERE tokenValue=$1 ORDER BY id ASC ", tokensTable)
	err := r.db.Get(&newToken, query, token)
	if newToken == (carsBrandsBattle.Token{}) {
		return &carsBrandsBattle.Token{}, nil
	}
	return &newToken, err
}

func (r *TokenPostgres) Update(token string, updatedToken carsBrandsBattle.UpdateTokenInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updatedToken.Revoked != nil {
		setValues = append(setValues, fmt.Sprintf("revoked=$%d", argId))
		args = append(args, *updatedToken.Revoked)
		argId++
	}

	if updatedToken.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("userId=$%d", argId))
		args = append(args, *updatedToken.UserId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s t SET %s WHERE t.tokenValue=$%d", tokensTable, setQuery, argId)
	args = append(args, token)
	_, err := r.db.Exec(query, args...)
	return err
}
