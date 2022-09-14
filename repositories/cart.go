package repositories

import (
	"server_wb/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCarts(UserID int) ([]models.Cart, error)
	GetCart(UserID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	CleaningCart(cart []models.Cart) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart

	err := r.db.Preload("User").Preload("Product").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCarts(UserID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Raw("SELECT * FROM carts WHERE user_id=?", UserID).Preload("Product").Preload("User").Scan(&carts).Error
	//err := r.db.Preload("User").Preload("Product").First(&carts, UserID).Error
	return carts, err
}

func (r *repository) GetCart(UserID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("Product").First(&cart, UserID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("User").Preload("Product").Create(&cart).Error // Using Create method

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error // Using Save method

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}

func (r *repository) CleaningCart(carts []models.Cart) ([]models.Cart, error) {
	err := r.db.Delete(&carts).Error
	return carts, err
}
