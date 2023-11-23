package service

import (
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
	"math/rand"
	"time"
)

type BrandService struct {
	repo repository.Brand
}

func NewBrandService(repo repository.Brand) *BrandService {
	return &BrandService{
		repo: repo,
	}
}

func (s *BrandService) Create(brand carsBrandsBattle.Brand) (int, error) {
	return s.repo.Create(brand)
}

func (s *BrandService) GetAll() ([]carsBrandsBattle.Brand, error) {
	return s.repo.GetAll()
}

func (s *BrandService) GetById(id int) (*carsBrandsBattle.Brand, error) {
	return s.repo.GetById(id)
}

func (s *BrandService) GetOneByRandom() (carsBrandsBattle.Brand, error) {
	w := rand.NewSource(time.Now().Unix())
	r := rand.New(w)
	brands, err := s.GetAll()
	value := r.Intn(len(brands))
	return brands[value], err
}

func (s *BrandService) Update(id int, brand carsBrandsBattle.UpdateBrandInput) error {
	if err := brand.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, brand)
}

func (s *BrandService) Delete(id int) error {
	return s.repo.Delete(id)
}
