package api

import (
	"crud-with-auth/db"
	"fmt"

	// "strings"

	"github.com/gin-gonic/gin"
)

func (api *Api) User() {
	group := api.r.Group("/user")
	group.Use(api.AuthMiddleware)
	group.GET("/profile", api.GetPofile)
}

// @Summary User profile
// @Accept  json
// @Produce  json
// @Tags     User
// @Param Authorization header string true "Authorization"
// @Router /user/profile [get]
// @Success 200 {object} db.User
// @Failure 404
func (api Api) GetPofile(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var user db.User
	if err := api.db.Storage.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, user)
}
