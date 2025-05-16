package dto

type RegisterInput struct {
	Name     string `json:"name" binding:"required" example:"Rian"`
	Email    string `json:"email" binding:"required,email" example:"ian@gmail.com"`
	Password string `json:"password" binding:"required,min=6" example:"P@ssW0rd"`
	//Role     string `json:"role" binding:"required" example:"admin / customer"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"ian@gmail.com"`
	Password string `json:"password" binding:"required" example:"P@ssW0rd"`
}
