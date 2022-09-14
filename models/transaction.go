package models

import "time"

type Transaction struct {
	ID        int             `json:"id"`
	UserID    int             `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserTransaction `json:"user"`
	Total     int             `json:"total" gorm:"type: int"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type TransactionResponse struct {
	ID     int             `json:"id"`
	UserID int             `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   UserTransaction `json:"user"`
	Total  int             `json:"total" gorm:"type: int"`
	Status string          `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
