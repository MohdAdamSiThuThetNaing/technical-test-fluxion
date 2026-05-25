package ai

import (
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
)

func RegisterRoutes(r *gin.Engine) {

	aiRoutes := r.Group("/ai")

	// Protected route
	aiRoutes.GET(
		"/test",
		guard.AuthMiddleware(),
		TestAI,
	)

	// Public route
	aiRoutes.POST(
		"/user-suggestion",
		SuggestUser,
	)
}