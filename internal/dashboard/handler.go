package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func Dashboard(c *gin.Context) {

	var totalUsers int64

	db.DB.Model(&users.User{}).
		Count(&totalUsers)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"TotalUsers": totalUsers,
	})
}