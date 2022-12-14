package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Stock     int       `json:"stock" gorm:"type: int"`
	Desc      string    `json:"desc" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductTransaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Desc  string `json:"desc" gorm:"type: varchar(255)"`
	Image string `json:"image"`
}

func (ProductTransaction) TableName() string {
	return "products"
}
