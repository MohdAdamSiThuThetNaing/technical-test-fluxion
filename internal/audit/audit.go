package audit

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/queue"
)

func PublishAuditLog(
	c *gin.Context,
	event Event,
	userID string,
	email string,
	name string,
	reason string,
) {

	queue.Publish(gin.H{
		"event":        string(event),
		"user_id":      userID,
		"user_email":   email,
		"updated_name": name,
		"reason":       reason,
		"ip_address":   c.ClientIP(),
		"user_agent":   c.Request.UserAgent(),
		"action_by":    email,
		"created_at":   time.Now(),
	})
}