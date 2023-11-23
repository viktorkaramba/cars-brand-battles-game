package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user carsBrandsBattle.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (*carsBrandsBattle.User, error) {
	var user carsBrandsBattle.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	if user == (carsBrandsBattle.User{}) {
		return &carsBrandsBattle.User{}, nil
	}
	return &user, err
}

func (r *AuthPostgres) GetUserByUsername(username string) (*carsBrandsBattle.User, error) {
	var user carsBrandsBattle.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)
	if user == (carsBrandsBattle.User{}) {
		return &carsBrandsBattle.User{}, nil
	}
	return &user, err
}
