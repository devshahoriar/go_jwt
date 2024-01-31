package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const JWT_Secret string = "JWT_Secret"

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("error to HashPassword")
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

type LogedUser struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GetJwt(id uint, email string) string {
	claims := LogedUser{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(JWT_Secret))

	return ss
}

func GetUserDataFromJWT(tokenString string) (*LogedUser, error) {
	if len(tokenString) == 0 {
		return nil, errors.New("token not found")
	}
	token, err := jwt.NewParser().ParseWithClaims(strings.Split(tokenString, " ")[1], &LogedUser{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_Secret), nil
	})
	if err != nil {
		return nil, err
	}
	d := token.Claims.(*LogedUser)

	return d, nil

}
