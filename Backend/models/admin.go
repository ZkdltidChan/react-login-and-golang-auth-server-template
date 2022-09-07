package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	// gorm.Model
	ID        string `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (b *Admin) TableName() string {
	return "admin"
}

func (x *Admin) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}
