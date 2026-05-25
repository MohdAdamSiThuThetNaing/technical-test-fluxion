package ai

import (
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
)

func RegisterRoutes(r *gin.Engine) {

	aiRoutes := r.Group("/ai")
	aiRoutes.Use(
		guard.AuthMiddleware(),
	)

	aiRoutes.GET(
		"/test",
		TestAI,
	)
}