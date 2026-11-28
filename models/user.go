package models

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required,max=100"`
	Surname string `json:"surname" binding:"required,max=100"`
	Email   string `json:"email" binding:"required,email,max=255"`
}
