package db

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// swagger:parameters getUserDetail
type Admin struct {
	Model
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GetAdminParam struct {
	Model
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

func (admin *Admin) GenerateAdminPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)
	if err != nil {
		panic(err)
	}
	admin.Password = string(bytes)
}

func (admin *Admin) ValidateAdminPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
}

type SignedAdminDetails struct {
	ID       uint
	Email    string
	UserName string
	jwt.StandardClaims
}

func (admin *Admin) AdminToken() (string, error) {
	claims := &SignedAdminDetails{
		ID:       admin.ID,
		Email:    admin.Email,
		UserName: admin.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))
}

func (admin *Admin) ValidateAdminToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedAdminDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*SignedAdminDetails); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return false, errors.New("token expired")
		}
		admin.ID = claims.ID
		admin.Email = claims.Email
		admin.Username = claims.UserName

		return true, nil
	} else {
		return false, errors.New("claim invalid")
	}
}
