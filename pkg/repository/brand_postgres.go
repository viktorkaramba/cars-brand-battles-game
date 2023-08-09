package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type BrandPostgres struct {
	db *sqlx.DB
}

func NewBrandPostgres(db *sqlx.DB) *BrandPostgres {
	return &BrandPostgres{db: db}
}

func (r *BrandPostgres) Create(brand carsBrandRandomGenerator.Brand) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, imageBrand) values ($1, $2) RETURNING id", brandsTable)
	row := r.db.QueryRow(query, brand.Name, brand.ImageBrand)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BrandPostgres) GetAll() ([]carsBrandRandomGenerator.Brand, error) {
	var brands []carsBrandRandomGenerator.Brand
	query := fmt.Sprintf("SELECT * FROM %s", brandsTable)
	err := r.db.Select(&brands, query)
	return brands, err
}

func (r *BrandPostgres) GetById(id int) (carsBrandRandomGenerator.Brand, error) {
	var brand carsBrandRandomGenerator.Brand
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", brandsTable)
	err := r.db.Select(&brand, query, id)
	return brand, err
}

func (r *BrandPostgres) Update(id int, brand carsBrandRandomGenerator.UpdateBrandInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, imageBrand=$2 WHERE id=$3", brandsTable)
	_, err := r.db.Exec(query, brand.Name, brand.ImageBrand, id)
	return err
}

func (r *BrandPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", brandsTable)
	_, err := r.db.Exec(query, id)
	return err
}
