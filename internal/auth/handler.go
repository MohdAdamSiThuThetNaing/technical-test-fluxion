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

type AuthReason string

const (

	USER_NOT_FOUND AuthReason = "USER_NOT_FOUND"

	INVALID_PASSWORD AuthReason = "INVALID_PASSWORD"

	UNKNOWN_USER_ID AuthReason = "UNKNOWN"
)

func ShowLogin(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"login.html",
		nil,
	)
}

func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	var user users.User

	err := db.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		audit.PublishAuditLog(
			c,
			audit.AUTH_LOGIN_FAILED,
			string(UNKNOWN_USER_ID),
			email,
			"",
			string(USER_NOT_FOUND),
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

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		audit.PublishAuditLog(
			c,
			audit.AUTH_LOGIN_FAILED,
			user.ID.String(),
			user.Email,
			user.Name,
			string(INVALID_PASSWORD),
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
	err = session.Save()

	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"login.html",
			gin.H{
				"Error": "Failed to save session",
			},
		)
		return
	}

	audit.PublishAuditLog(
		c,
		audit.AUTH_LOGIN_SUCCESS,
		user.ID.String(),
		user.Email,
		user.Name,
		"",
	)

	c.Redirect(
		http.StatusFound,
		"/",
	)
}

func Logout(c *gin.Context) {

	session := sessions.Default(c)

	userID := session.Get("user_id")
	userEmail := session.Get("user_email")
	userName := session.Get("user_name")

	audit.PublishAuditLog(
		c,
		audit.AUTH_LOGOUT,
		toString(userID),
		toString(userEmail),
		toString(userName),
		"",
	)

	session.Clear()
	err := session.Save()

	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"login.html",
			gin.H{
				"Error": "Failed to logout",
			},
		)
		return
	}

	c.Redirect(
		http.StatusFound,
		"/login",
	)
}

func toString(value interface{}) string {

	if value == nil {
		return ""
	}

	str, ok := value.(string)
	if !ok {
		return ""
	}

	return str
}