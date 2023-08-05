package repository

type Brand interface {
}

type Repository struct {
	Brand
}

func NewRepository() *Repository {
	return &Repository{}
}
