package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func ShowLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	var user users.User
	err := db.DB.
		Where("email = ?", email).
		First(&user).Error

	// User not found
	if err != nil {
		audit.PublishAuditLog(
			c,
			"AUTH_LOGIN_FAILED",
			"",
			email,
			"",
			"USER_NOT_FOUND",
		)

		c.HTML(
			http.StatusUnauthorized,
			"login.html",
			gin.H{
				"Error": "Invalid credentials",
			},
		)

		return
	}

	// Invalid password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		audit.PublishAuditLog(
			c,
			"AUTH_LOGIN_FAILED",
			user.ID.String(),
			user.Email,
			user.Name,
			"INVALID_PASSWORD",
		)

		c.HTML(
			http.StatusUnauthorized,
			"login.html",
			gin.H{
				"Error": "Invalid credentials",
			},
		)

		return
	}

	session := sessions.Default(c)

	session.Set("authenticated", true)
	session.Set("user_id", user.ID.String())
	session.Set("user_email", user.Email)
	session.Set("user_name", user.Name)
	session.Save()
	audit.PublishAuditLog(
		c,
		"AUTH_LOGIN_SUCCESS",
		user.ID.String(),
		user.Email,
		user.Name,
		"",
	)
	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}