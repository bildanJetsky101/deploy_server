package profiledto

type CreateProfile struct {
	ID      int    `json:"id"`
	Phone   int    `json:"phone" gorm:"type: int"`
	Image   string `json:"image" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
	Gender  string `json:"gender" gorm:"type: varchar(255)"`
	UserID  int    `json:"user_id"`
}

type ProfileResponse struct {
	ID      int    `json:"id"`
	Phone   int    `json:"phone" gorm:"type: int"`
	Image   string `json:"image" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
	Gender  string `json:"gender" gorm:"type: varchar(255)"`
	UserID  int    `json:"user_id"`
}
