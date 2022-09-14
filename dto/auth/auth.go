package authdto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	IsAdmin  bool   `json:"is_Admin" form:"is_Admin" `
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}

type LoginResponse struct {
	Name    string `gorm:"type: varchar(255)" json:"name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	Token   string `gorm:"type: varchar(255)" json:"token"`
	IsAdmin bool   `json:"is_Admin" form:"is_Admin" `
}

type CheckAuthResponse struct {
	Id      int    `gorm:"type: int" json:"id"`
	Name    string `gorm:"type: varchar(255)" json:"name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	Token   string `gorm:"type: varchar(255)" json:"token"`
	IsAdmin bool   `json:"is_Admin" form:"is_Admin" `
}
