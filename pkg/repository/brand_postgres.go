package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"strings"
)

type BrandPostgres struct {
	db *sqlx.DB
}

func NewBrandPostgres(db *sqlx.DB) *BrandPostgres {
	return &BrandPostgres{db: db}
}

func (r *BrandPostgres) Create(brand carsBrandsBattle.Brand) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, imageBrand) values ($1, $2) RETURNING id", brandsTable)
	row := r.db.QueryRow(query, brand.Name, brand.ImageBrand)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BrandPostgres) GetAll() ([]carsBrandsBattle.Brand, error) {
	var brands []carsBrandsBattle.Brand
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC ", brandsTable)
	err := r.db.Select(&brands, query)
	return brands, err
}

func (r *BrandPostgres) GetById(id int) (*carsBrandsBattle.Brand, error) {
	var brand carsBrandsBattle.Brand
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1 ORDER BY id ASC ", brandsTable)
	err := r.db.Get(&brand, query, id)
	if brand == (carsBrandsBattle.Brand{}) {
		return &carsBrandsBattle.Brand{}, nil
	}
	return &brand, err
}

func (r *BrandPostgres) Update(id int, brand carsBrandsBattle.UpdateBrandInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if brand.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *brand.Name)
		argId++
	}

	if brand.ImageBrand != nil {
		setValues = append(setValues, fmt.Sprintf("imageBrand=$%d", argId))
		args = append(args, *brand.ImageBrand)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s b SET %s WHERE b.id=$%d", brandsTable, setQuery, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BrandPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", brandsTable)
	_, err := r.db.Exec(query, id)
	return err
}
