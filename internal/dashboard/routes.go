package dashboard

import (
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
)

func RegisterRoutes(r *gin.Engine) {

	dashboard := r.Group("/")
	dashboard.Use(guard.AuthMiddleware())
	dashboard.GET("", Dashboard)
}