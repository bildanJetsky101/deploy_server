package models

import "time"

type Cart struct {
	ID        int                `json:"id" gorm:"primary_key:auto_increment"`
	QTY       int                `json:"qty"`
	SubTotal  int                `json:"subtotal"`
	ProductID int                `json:"product_id" form:"product_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   ProductTransaction `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status    string             `json:"status"`
	UserID    int                `json:"user_id" form:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserProfile        `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time          `json:"-"`
	UpdatedAt time.Time          `json:"-"`
}

type CartResponse struct {
	ID        int                `json:"id"`
	QTY       int                `json:"qty"`
	Total     int                `json:"total"`
	SubTotal  int                `json:"subtotal"`
	ProductID int                `json:"product_id" form:"product_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   ProductTransaction `json:"product" gorm:"many2many:cart_products"`
	Status    string             `json:"status"`
	UserID    int                `json:"user_id" form:"user_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserProfile        `json:"user"`
}

type CartResponseUser struct {
	ID        int                `json:"id"`
	QTY       int                `json:"qty"`
	Total     int                `json:"total"`
	SubTotal  int                `json:"subtotal"`
	ProductID int                `json:"product_id" form:"product_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   ProductTransaction `json:"product" gorm:"many2many:cart_products"`
	Status    string             `json:"status"`
	UserID    int                `json:"user_id" form:"user_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CartResponseTransaction struct {
	ID        int                `json:"id"`
	QTY       int                `json:"qty"`
	Total     int                `json:"total"`
	SubTotal  int                `json:"subtotal"`
	ProductID int                `json:"product_id" form:"product_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   ProductTransaction `json:"product" gorm:"many2many:cart_products"`
	Status    string             `json:"status"`
	UserID    int                `json:"user_id" form:"user_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CartResponse) TableName() string {
	return "carts"
}

func (CartResponseUser) TableName() string {
	return "carts"
}

func (CartResponseTransaction) TableName() string {
	return "carts"
}
