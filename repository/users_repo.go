package repository

import (
	"wolftagon/models"

	"gorm.io/gorm"
)

type UserRepo interface{
	CreateUser(CreateUserReq models.CreateUserRequest) (*models.CreateUserResponse,error)
	GetUserEmail (email string)(*models.User,error)

} 

type userRepo struct{
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo{
	return &userRepo{
		db:db,
	}
}
