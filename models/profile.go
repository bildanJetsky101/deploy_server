package models

import "time"

type Profile struct {
	ID        int         `json:"id"`
	Phone     int         `json:"phone" gorm:"type: int"`
	Image     string      `json:"image" gorm:"type: varchar(255)"`
	Address   string      `json:"address" gorm:"type: varchar(255)"`
	UserID    int         `json:"user_id"`
	Gender    string      `json:"gender" gorm:"type: varchar(255)"`
	User      UserProfile `json:"user"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}

type ProfileResponse struct {
	ID      int         `json:"id"`
	Phone   int         `json:"phone"`
	Gender  string      `json:"gender"`
	Address string      `json:"address"`
	Image   string      `json:"image" gorm:"type: varchar(255)"`
	UserID  int         `json:"user_id"`
	User    UserProfile `json:"user"`
}

type ProfileTransaction struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}

func (ProfileTransaction) TableName() string {
	return "profiles"
}
