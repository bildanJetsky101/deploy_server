package models

import "time"

// User model struct
type User struct {
	ID            int                `json:"id"`
	Name          string             `json:"name" gorm:"type: varchar(255)"`
	Email         string             `json:"email" gorm:"type: varchar(255)"`
	IsAdmin       bool               `json:"is_Admin" gorm:"type: boolean"`
	Password      string             `json:"password" gorm:"type: varchar(255)"`
	Profile       ProfileResponse    `json:"profile"`
	Cart          []CartResponseUser `json:"cart"`
	TransactionID int                `json:"transaction_id"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

type UserResponse struct {
	ID        int                `json:"id"`
	Name      string             `json:"name" gorm:"type: varchar(255)"`
	Email     string             `json:"email" gorm:"type: varchar(255)"`
	IsAdmin   bool               `json:"is_Admin" gorm:"type: boolean"`
	Password  string             `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileResponse    `json:"profile"`
	Cart      []CartResponseUser `json:"cart" `
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UserProfile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserTransaction struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	ProfileID     int                `json:"-"`
	TransactionID int                `json:"-"`
	Profile       ProfileTransaction `json:"profile" gorm:"many2many:user_profiles"`
	CartID        int                `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//Cart          []CartResponseUser `json:"cart" `
}

func (UserProfile) TableName() string {
	return "users"
}

func (UserTransaction) TableName() string {
	return "users"
}

func (UserResponse) TableName() string {
	return "users"
}
