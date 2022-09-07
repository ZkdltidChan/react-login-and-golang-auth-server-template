package controllers

import (
	"crud-golang/middlewares"
	"crud-golang/models"
	"crud-golang/pkg/e"
	// "crud-golang/pkg/response"
	"crud-golang/pkg/utils"
	"crud-golang/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	// _ "github.com/go-sql-driver/mysql"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func AuthHandler(c *gin.Context) {
	var code int
	var user models.User
	// var res response.ResponseMessage

	if binderr := c.ShouldBindJSON(&user); binderr != nil {
		log.Error().Err(binderr).
			Msg("Error occurred while binding request data")
		// response.SendResponse(res.)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}
	vu, err := repo.Login(&user)
	if err != nil {
		code = e.ERROR_AUTH
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": nil,
		})
	}

	if vu == nil {
		code = e.ERROR_USER_NOT_FOUND
		c.JSON(http.StatusNotFound, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": nil,
		})
	}

	if vu != nil {
		tokenString, _ := middlewares.CreateToken(vu.Username)
		code = e.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": gin.H{"token": tokenString, "username": vu.Username},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Verified Failed.",
	})
	return
}

func CreateUser(c *gin.Context) {
	var user models.User
	if bindErr := c.ShouldBindJSON(&user); bindErr != nil {
		log.Error().Err(bindErr).
			Msg("Error occurred while binding request data")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": bindErr.Error(),
		})
		return
	}
	user.FillDefaults()

	hashedPassword, hashErr := Hash(user.Password)
	if hashErr != nil {
		code := e.ERROR_HASH_PASSWORD
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	user.Password = string(hashedPassword)

	userInfo, err := repo.CreateUser(&user)
	if err != nil {
		code := e.ERROR_UNDIFINE
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userInfo,
	})
}

func GetAllUsers(c *gin.Context) {
	var user []models.User
	userDB := repo.GetAllUsers(&user)
	userData := utils.PageOperation(c, userDB, &user)
	c.JSON(http.StatusOK, gin.H{
		"data": userData,
	})

}

// func GetUserById(c *gin.Context) {

// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func UpdateUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		println("no data found")
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	c.BindJSON(&user)
// 	err = Model.UpdateUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func DeleteUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.DeleteUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// }
