package response

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, res ResponseMessage) {
	if len(res.Message) > 0 {
		c.JSON(res.Status, map[string]interface{}{"message": strings.Join(res.Message, "; ")})
	} else if len(res.Error) > 0 {
		c.JSON(res.Status, map[string]interface{}{"error": strings.Join(res.Error, "; ")})
	}
}
