package auth

import (
	"errors"
	"fmt"
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

func VerifyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); ok {
			return JWT_PUBLIC_KEY, nil
		}

		return nil, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("malformed token")
		}
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, fmt.Errorf("token expired or not valid yet")
		}

		return nil, fmt.Errorf("token invalid")
	}

	return token, nil
}