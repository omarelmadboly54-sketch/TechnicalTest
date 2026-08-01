package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB,error){
	_=godotenv.Load()
	dsn:=os.Getenv("DB_URL")
	if dsn ==""{
		return nil,fmt.Errorf("Db_Url Is Not Set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil{
		return nil,err
	}
	return db,nil

}