package db

import (
	"gorm.io/gorm"
	"os"
	"time"
)

var SECRET_KEY string = os.Getenv("API_SECRET")

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
