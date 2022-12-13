package api

import (
	"crud-with-auth/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) Auth() {
	noAuth := api.r.Group("/user")
	noAuth.POST("/login", api.UserLoginHandler)
	noAuth.POST("/register", api.UserRegisterHandler)
	noAuth.POST("/refresh-token", api.UserRefreshHandler)
}

// @Summary User login
// @Accept  json
// @Produce  json
// @Tags     User
// @Param data body db.LoginUser true "body data".
// @Router /user/login [post]
// @Success 200 {object} db.User
// @Failure 400
func (api Api) UserLoginHandler(c *gin.Context) {
	var userDto db.User
	var user db.User
	c.BindJSON(&userDto)

	if err := api.db.Storage.Where("username = ?", userDto.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Username or password invalid",
		})
		return
	}
	if err := user.ValidatePassword(userDto.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Username or password invalid",
		})
		return
	}

	if token, err := user.Token(); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"user":    user,
			"token":   token,
		})
	}
}

// @Summary Create User
// @Accept  json
// @Produce  json
// @Tags     User
// @Param data body db.RegisterUser true "body data".
// @Router /user/register [post]
// @Success 201 {object} db.User
// @Failure 400
func (api Api) UserRegisterHandler(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)
	user.GeneratePassword()
	if err := api.db.Storage.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered",
		"username": user.Username,
	})
}

func (api Api) UserRefreshHandler(c *gin.Context) {

}
