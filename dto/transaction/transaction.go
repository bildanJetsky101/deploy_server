package transactiondto

type CreateTransaction struct {
	UserID int    `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total  int    `json:"total" gorm:"type: int"`
	Status string `json:"status"`
}

type ResponseTransaction struct {
	UserID int    `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total  int    `json:"total" gorm:"type: int"`
	Status string `json:"status"`
}
