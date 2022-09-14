package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	IsAdmin  bool   `json:"is_Admin" form:"is_Admin" `
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	IsAdmin  bool   `json:"is_Admin" form:"is_Admin" `
	Password string `json:"password" form:"password"`
}
