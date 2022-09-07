package repo

import (
	// "crud-golang/config"
	"crud-golang/models"
	"crud-golang/pkg/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) (*models.User, error) {
	db, dbConErr := utils.GetDatabaseConnection()
	if dbConErr != nil {
		log.Err(dbConErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, dbConErr
	}

	result := db.Create(&user)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return user, nil
}

func GetAllUsers(user *[]models.User) (db *gorm.DB) {
	db, dbConErr := utils.GetDatabaseConnection()
	db = db.Model(&user)
	if dbConErr != nil {
		log.Err(dbConErr).Msg("Error occurred while getting a DB connection from the connection pool")
	}
	return
}
