package repo

import (
	// "crud-golang/config"
	"crud-golang/models"
	"crud-golang/pkg/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	db, dbConErr := utils.GetDatabaseConnection()
	if dbConErr != nil {
		log.Err(dbConErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, dbConErr
	}

	result := db.Create(&admin)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return admin, nil
}

func GetAllAdmins(admin *[]models.Admin) (db *gorm.DB) {
	db, dbConErr := utils.GetDatabaseConnection()
	db = db.Model(&admin)
	if dbConErr != nil {
		log.Err(dbConErr).Msg("Error occurred while getting a DB connection from the connection pool")
	}
	return
}
