package repo

import (
	"crud-golang/models"
	"crud-golang/pkg/utils"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(user *models.User) (*models.User, error) {
	var uu models.User
	db, conErr := utils.GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, conErr
	}

	var err error
	err = db.Model(user).Where("email = ?", user.Email).Take(&uu).Error
	if err != nil {
		return nil, err
	}
	err = VerifyPassword(uu.Password, user.Password)
	if err != nil {
		return nil, err
	}
	return &uu, nil
}

func AdminLogin(admin *models.Admin) (*models.Admin, error) {
	var uu models.Admin
	db, conErr := utils.GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, conErr
	}

	var err error
	err = db.Model(admin).Where("username = ?", admin.Username).Take(&uu).Error
	if err != nil {
		return nil, err
	}
	err = VerifyPassword(uu.Password, admin.Password)
	if err != nil {
		return nil, err
	}
	return &uu, nil
}
