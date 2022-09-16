package repositories

import (
	"server_wb/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransactions(UserID int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	//GetCartsTransaction(UserID int) ([]models.Cart, error)
	CreateTransaction(cart models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetTransactions(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Raw("SELECT * FROM transactions WHERE user_id=?", UserID).Preload("Product").Scan(&transactions).Error
	//err := r.db.Preload("User").Preload("Product").First(&carts, UserID).Error
	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.First(&transactions, ID).Error

	return transactions, err
}

// func (r *repository) GetCartsTransaction(UserID int) ([]models.Cart, error) {
// 	var carts []models.Cart
// 	err := r.db.Raw("SELECT * FROM carts WHERE user_id=?", UserID).Preload("Product").Scan(&carts).Error
// 	//err := r.db.Preload("User").Preload("Product").First(&carts, UserID).Error
// 	return carts, err
// }

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&transaction).Error // Using Create method

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Toping").First(&transaction, ID)

	if status != transaction.Status && status == "success" {

	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Product").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
