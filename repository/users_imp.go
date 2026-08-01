package repository

import "wolftagon/models"

func (u *userRepo) CreateUser(CreateUserReq models.CreateUserRequest) (*models.CreateUserResponse,error){
	newUser:=models.User{
		Name: CreateUserReq.Name,
		Email: CreateUserReq.Email,
		Password: CreateUserReq.Password,
	}
	err:=u.db.Create(&newUser).Error
	if err!=nil{
		return nil,err
	}
	response:=&models.CreateUserResponse{
		Name: newUser.Name,
		Email: newUser.Email,
	}
	return response,nil
}


func(u *userRepo) GetUserEmail (email string)(*models.User,error){
	var user models.User

	err:=u.db.Where("email=?",email).First(&user).Error
	if err!=nil{
		return nil,err
	}
	return &user,nil
}
