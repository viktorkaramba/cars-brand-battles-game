package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
