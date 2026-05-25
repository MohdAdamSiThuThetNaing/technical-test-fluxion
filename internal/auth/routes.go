package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	r.GET("/login", ShowLogin)
	r.POST("/login", Login)
	r.GET("/logout", Logout)
}