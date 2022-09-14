package productdto

type CreateProduct struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Stock int    `json:"stock" form:"stock" gorm:"type: int" validate:"required"`
	Desc  string `json:"desc" form:"desc" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Stock int    `json:"stock" form:"stock" gorm:"type: int" validate:"required"`
	Desc  string `json:"desc" form:"desc" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}
