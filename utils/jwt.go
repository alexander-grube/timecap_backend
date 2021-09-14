package utils

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}