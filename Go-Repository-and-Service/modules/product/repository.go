package product

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product *Product) error
	FindAll() ([]Product, error)
	FindOne(id int) (Product, error)
	Update(product *Product) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Repository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindOne(id int) (Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *productRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id int) error {
	result := r.db.Delete(&Product{}, id)
	if result.RowsAffected == 0 {
		return errors.New("Product not found")
	}
	return result.Error
}
