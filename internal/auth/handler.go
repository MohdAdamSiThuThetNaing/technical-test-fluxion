package auth

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if email != adminEmail || password != adminPassword {

		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Invalid credentials",
		})
		return
	}

	session := sessions.Default(c)

	session.Set("admin_logged_in", true)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}