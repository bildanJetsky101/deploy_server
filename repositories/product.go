package repositories

import (
	"server_wb/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Raw("SELECT * FROM products").Scan(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Raw("SELECT * FROM products WHERE id=?", ID).Scan(&product).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error // Using Create method

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error // Using Save method

	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
