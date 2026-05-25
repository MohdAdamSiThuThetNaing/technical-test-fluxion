package logs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {

	service := NewService()
	logs, err := service.GetLogs()
	
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"logs.html",
		gin.H{
			"Logs": logs,
		},
	)
}