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
	// _ "github.com/go-sql-driver/mysql"
)

func AdminAuthHandler(c *gin.Context) {
	var code int
	var admin models.Admin
	// var res response.ResponseMessage

	if binderr := c.ShouldBindJSON(&admin); binderr != nil {
		log.Error().Err(binderr).
			Msg("Error occurred while binding request data")
		// response.SendResponse(res.)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}
	vu, err := repo.AdminLogin(&admin)
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
			"code":         code,
			"msg":          e.GetMsg(code),
			"access_token": tokenString,
			"data":         gin.H{"name": vu.Username},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Verified Failed.",
	})
	return
}

func CreateAdmin(c *gin.Context) {
	var admin models.Admin
	if bindErr := c.ShouldBindJSON(&admin); bindErr != nil {
		log.Error().Err(bindErr).
			Msg("Error occurred while binding request data")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": bindErr.Error(),
		})
		return
	}
	admin.FillDefaults()

	hashedPassword, hashErr := Hash(admin.Password)
	if hashErr != nil {
		code := e.ERROR_HASH_PASSWORD
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	admin.Password = string(hashedPassword)

	adminInfo, err := repo.CreateAdmin(&admin)
	if err != nil {
		code := e.ERROR_UNDIFINE
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": adminInfo,
	})
}

func GetAllAdmins(c *gin.Context) {
	var admin []models.Admin
	adminDB := repo.GetAllAdmins(&admin)
	adminData := utils.PageOperation(c, adminDB, &admin)
	c.JSON(http.StatusOK, gin.H{
		"data": adminData,
	})

}
