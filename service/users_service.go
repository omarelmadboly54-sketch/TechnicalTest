package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"wolftagon/models"
	"wolftagon/repository"
	"wolftagon/utils"
)

type UserService interface{
	CreateUser (req models.CreateUserRequest)(*models.CreateUserResponse,error)
	LoginUser (req models.LoginRequest)(models.LoginResponse,error)
	FetchProducts()([]models.Product,error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService{
	return &userService{
		userRepo:userRepo,
	}
}


func (s *userService) CreateUser (req models.CreateUserRequest)(*models.CreateUserResponse,error){
	existingUser,_:=s.userRepo.GetUserEmail(req.Email)
	if existingUser!=nil{
		return nil,errors.New("User with This Email Already Exist")
	}
	hashedpassword,err:=utils.HashPassword(req.Password)
	if err!= nil{
		return nil,errors.New("failed to hash password")
	}
	hashedreq:=req
	hashedreq.Password=hashedpassword
	response,err:=s.userRepo.CreateUser(hashedreq)
	if err!=nil{
		return nil,err
	}
	return response,nil


}

func(s *userService) LoginUser (req models.LoginRequest)(models.LoginResponse,error){
	user,err:=s.userRepo.GetUserEmail(req.Email)
	if err!= nil{
		return models.LoginResponse{},errors.New("invalid credentials")
	}
	if ! utils.CheckPasswordHash(req.Password,user.Password){
		return models.LoginResponse{},errors.New("invalid credentials")
	}
	tokenstring,err:=utils.GenerateToken(*user)
	if err!=nil{
		return models.LoginResponse{},err
	}
	return models.LoginResponse{Token:tokenstring},nil
}


type DummyJSONResponse struct{
	Products 	[]models.Product	`json:"products"`
}

func (s *userService)FetchProducts()([]models.Product,error){
	resp,err:=http.Get("https://dummyjson.com/products?utm_source=chatgpt.com")
	if err!= nil{
		return nil,err
	}
	defer resp.Body.Close()

	var result DummyJSONResponse
	if err:=json.NewDecoder(resp.Body).Decode(&result);err!=nil{
		return nil,err
	}
	return result.Products,nil
}