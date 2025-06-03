package models

type AuthStruct struct {
	Id       int    `json:"id" form:"id" db:"id"`
	Email    string `json:"email" form:"email" db:"email" binding:"required,email" `
	Password string `json:"password,omitempty" form:"password" db:"password" binding:"required,min=6"`
	Role     string `json:"role,omitempty" form:"role" db:"role"`
}
