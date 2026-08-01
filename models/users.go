package models

import "time"

type User struct {
	ID        	uint   		`gorm:"primarykey" json:"id"`
	Name      	string 		`gorm:"not null" json:"name"`
	Email     	string 		`gorm:"unique;not null" json:"email"`
	Password  	string 		`gorm:"not null" json:"password"`
	CreatedAt 	time.Time	`gorm:"autoCreateTime" json:"created_at"`
}

type CreateUserRequest struct {
	Name		string		`json:"name" binding:"required"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct{
	Name		string		`json:"name"`
	Email		string		`json:"email"`
}

type LoginRequest struct{
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required"`
}

type LoginResponse struct{
	Token		string		`json:"token"`
}
