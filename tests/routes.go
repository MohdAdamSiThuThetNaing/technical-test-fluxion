package tests

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	r.GET(
		"/tests",
		ShowTests,
	)
}