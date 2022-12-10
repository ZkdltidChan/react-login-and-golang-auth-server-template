package api

import (
	"crud-with-auth/db"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
)

func (api *Api) Admin() {
	group := api.r.Group("/admin")
	group.POST("/login", api.AdminLoginHandler)
	group.POST("/register", api.AdminRegisterHandler)

	group.Use(api.AuthAdminMiddleware)
	group.GET("/users", api.GetUserList)
}

// @Summary User List
// @Accept  json
// @Produce  json
// @Tags     Admin
// @Param Authorization header string true "Authorization"
// @Router /admin/users [get]
// @Success 200 {object} []db.Admin
// @Failure 400
func (api Api) GetUserList(c *gin.Context) {
	admin_id, _ := c.Get("admin_id")
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiRW1haWwiOiIiLCJVc2VyTmFtZSI6ImNoYW5fYWRtaW4iLCJleHAiOjE2NzA2MTY5ODh9.n6QVJi8GEVO-XWfxqDxmzlm2zjtzhkDVAs5McLgztqs
	var admin db.Admin
	var users []db.User
	if err := api.db.Storage.Where("id = ?", admin_id).First(&admin).Error; err != nil {
		c.AbortWithStatus(401)
		return
	}
	if err := api.db.Storage.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, users)
}

// @Summary Admin login
// @Accept  json
// @Produce  json
// @Tags     Admin
// @Param data body db.LoginAdmin true "body data".
// @Router /admin/login [post]
// @Success 200 {object} db.Admin
// @Failure 400
func (api Api) AdminLoginHandler(c *gin.Context) {
	var adminDto db.Admin
	var admin db.Admin
	c.BindJSON(&adminDto)

	if err := api.db.Storage.Where("username = ?", adminDto.Username).First(&admin).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Username or password invalid",
		})
		return
	}

	if err := admin.ValidateAdminPassword(adminDto.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Username or password invalid",
		})
		return
	}

	if token, err := admin.AdminToken(); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Admin Login successfully",
			"token":   token,
		})
	}
}

// @Summary Create Admin
// @Accept  json
// @Produce  json
// @Tags     Admin
// @Param data body db.RegisterAdmin true "body data".
// @Router /admin/register [post]
// @Success 201 {object} db.Admin
// @Failure 400
func (api Api) AdminRegisterHandler(c *gin.Context) {
	var admin db.Admin
	c.BindJSON(&admin)
	admin.GenerateAdminPassword()
	if err := api.db.Storage.Create(&admin).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin registered",
	})
}
