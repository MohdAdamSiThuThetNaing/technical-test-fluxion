package logs

import (
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
)

func RegisterRoutes(r *gin.Engine) {

	logRoutes := r.Group("/logs")
	logRoutes.Use(
		guard.AuthMiddleware(),
	)

	logRoutes.GET(
		"",
		GetLogs,
	)
}