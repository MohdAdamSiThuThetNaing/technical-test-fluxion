package users

import (
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/guard"
)

func RegisterRoutes(r *gin.Engine) {

	users := r.Group("/users")

	users.Use(guard.AuthMiddleware())

	users.GET("", ListUsers)
	users.GET("/create", ShowCreateUser)
	users.POST("/create", CreateUser)

	users.GET("/edit/:id", ShowEditUser)
	users.POST("/edit/:id", UpdateUser)

	users.POST("/delete/:id", DeleteUser)
}