package db

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GetUserParam struct {
	Model
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

func (user *User) GeneratePassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		panic(err)
	}
	user.Password = string(bytes)
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

type SignedDetails struct {
	ID       uint
	Email    string
	UserName string
	jwt.StandardClaims
}

func (user *User) Token() (string, error) {
	claims := &SignedDetails{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))
}

func (user *User) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*SignedDetails); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return false, errors.New("token expired")
		}
		user.ID = claims.ID
		user.Email = claims.Email
		user.Username = claims.UserName

		return true, nil
	} else {
		return false, errors.New("claim invalid")
	}
}
