package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_PRIVATE_KEY = GetJwtKey("private.pem")
	JWT_PUBLIC_KEY  = GetJwtKey("public.pem")
)

func GetJwtKey(key string) interface{} {
	value, err := os.ReadFile(key)

	if err != nil {
		log.Fatal(err)
	}

	var rsaKey interface{}

	switch key {
	case "private.pem":
		rsaKey, err = jwt.ParseRSAPrivateKeyFromPEM(value)
	case "public.pem":
		rsaKey, err = jwt.ParseRSAPublicKeyFromPEM(value)
	}

	if err != nil {
		log.Fatal(err)
	}
	
	return rsaKey
}

func GenerateToken(identifier string) (string, error){
	claims := jwt.MapClaims{
		"sub": identifier,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	
	return token.SignedString(JWT_PRIVATE_KEY)
}