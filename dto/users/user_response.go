package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	IsAdmin  bool   `json:"is_Admin" form:"is_Admin" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
