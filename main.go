package main

import (
	"log"
	"net/http"
	"wolftagon/configuration"
	"wolftagon/controllers"
	"wolftagon/models"
	"wolftagon/repository"
	"wolftagon/service"

	"github.com/gin-gonic/gin"
)

func main(){
	db, err:= configuration.InitPostgres()
	if err!= nil{
		panic("Failed To Connect To database"+err.Error())
	}
	log.Println("DataBase Connected Successfully")
	err=db.AutoMigrate(models.User{})
	if err!=nil{
		panic("Failed to Migrate DataBase"+err.Error())
	}
	log.Println("DataBase Migration Completed Successfully")
	userRepo:=repository.NewUserRepo(db)
	userService:=service.NewUserService(userRepo)
	userController:=controllers.NewUserController(userService)


	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("signup",func(c *gin.Context){
		c.HTML(http.StatusOK,"signup.html",nil)
	})
	r.GET("signin",func(c *gin.Context) {
		c.HTML(http.StatusOK,"signin.html",nil)
	})
	r.GET("/",userController.GetProducts)
	r.GET("/checkout",func(c *gin.Context){
		token,err:=c.Cookie("token")
		if err!=nil || token ==""{
			c.Redirect(http.StatusSeeOther,"signin")
			return
		}
		c.HTML(http.StatusOK,"checkout.html",nil)
	}) 

	r.POST("/api/signup",userController.CreateUser)
	r.POST("/api/signin",userController.SignIn)

	log.Println("server is running on http://127.0.0.1:8080")
	if err:=r.Run(":8080");err!=nil{
		panic("failed to start server"+err.Error())
	}


}