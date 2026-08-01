package controllers

import (
	"net/http"
	"wolftagon/models"
	"wolftagon/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController{
	return &userController{
		userService: userService,
	}
}

func (uc *userController) CreateUser(c *gin.Context){
	var req models.CreateUserRequest

	if err:=c.ShouldBind(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request payload"+err.Error(),
		})
		return
	}
	res,err:=uc.userService.CreateUser(req)
	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"User Registered Successfully",
		"data":res,
	})
}

func (uc *userController) SignIn(c *gin.Context){
	var req models.LoginRequest

	if err:=c.ShouldBind(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request payload"+err.Error(),
		})
		return
	}
		res,err:=uc.userService.LoginUser(req)
	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.SetCookie("token",res.Token,3600,"/","",false,true)
	
	c.JSON(http.StatusCreated,gin.H{
		"message":"Login Successfully",
		"data":res,
	})


}

func (uc *userController) GetProducts (c *gin.Context){
	products,err:=uc.userService.FetchProducts()
	if err!=nil{
		c.HTML(http.StatusInternalServerError,"error.html",gin.H{
			"error":"failed to fetch products",
		})
		return
	}
	c.HTML(http.StatusOK,"products.html",gin.H{
		"products":products,
	})
}