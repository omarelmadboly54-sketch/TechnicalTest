package utils

import (
	"errors"
	"os"
	"time"
	"wolftagon/models"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user models.User) (string,error){
	secret:=os.Getenv("Jwt_SECRET")
	if secret ==""{
		return "",errors.New("JWT_SECRET Is Not Set In Env")
	}
	claims:=jwt.MapClaims{
		 "user_id": user.ID, 
        	"exp": time.Now().Add(time.Hour * 24).Unix(), 
        }
		token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
		tokenString, err := token.SignedString([]byte (secret))
    	if err != nil {
    	return "", err
    }

 		return tokenString, nil
}
	
