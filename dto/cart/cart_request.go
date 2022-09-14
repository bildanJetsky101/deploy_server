package cartdto

type CreateCart struct {
	ProductID int    `json:"product_id"`
	QTY       int    `json:"qty" form:"price" gorm:"type: int" validate:"required"`
	SubTotal  int    `json:"subtotal" form:"subtotal" gorm:"type: int" validate:"required"`
	Status    string `jsom:"status" form:"status" validate:"required"`
}

type UpdateCart struct {
	ProductID int `json:"product_id"`
}
