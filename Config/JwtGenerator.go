package config

import (
	"log"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
		Email string `json:"EmailId"`
		jwt.RegisteredClaims
}

func GenerateAuthToken(emailId string)(string,error){

	secretKey:= []byte("7c8f2a91e5b44d4f9a2c7e8d1f6b3a4c9d8e7f6a5b4c3d2e1f0a9b8c7d6e5f4")

	claims:= &Claims{
		Email: emailId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute*5)),
		},
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	SignedToken, err:= token.SignedString(secretKey)
	if err!= nil{
		log.Printf("The token is not signed with secret key... %v",err)
		return "",err
	}
	return SignedToken,nil
}