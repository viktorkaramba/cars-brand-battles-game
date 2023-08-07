package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Brand interface {
}

type Battle interface {
}

type Score interface {
}

type Repository struct {
	Authorization
	Brand
	Battle
	Score
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
