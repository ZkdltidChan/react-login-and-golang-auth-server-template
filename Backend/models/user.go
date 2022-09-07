package models

import (
	"github.com/google/uuid"
)

type User struct {
	// gorm.Model
	ID        string `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password" `
	Email     string `json:"email"`
	NickName  string `json:"nick_name"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// type UserResponse struct {
// 	User
// 	Password struct{} `json:"-"`
// }
// type LoginUser struct {
// 	Username string `json:"username"`
// 	Password string `json:"password" `
// }

func (b *User) TableName() string {
	return "user"
}

func (x *User) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}

// func (x *User) SPassword() {
// 	x.Password = ""
// }
