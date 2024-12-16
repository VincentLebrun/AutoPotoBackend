package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,alpha"`
	Surname  string `json:"surname" validate:"required,alpha"`
	Age      int    `json:"age" validate:"required,gte=18,lte=100"`
	Phone    string `json:"phone" validate:"required,phoneFR"`
	Mail     string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
	Role     string `json:"role" validate:"required,alpha"`
}
