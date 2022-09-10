package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type ProviderDB struct {
	Storage *gorm.DB
}

func GetDSN() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := "postgres"
	port := os.Getenv("POSTGRES_PORT")
	ssl := os.Getenv("POSTGRES_SSL")
	timezone := os.Getenv("POSTGRES_TIMEZONE")
	// fmt.Printf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
}

func ProvideDB() ProviderDB {
	// db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})
	db.AutoMigrate(&User{})

	return ProviderDB{Storage: db}
}
