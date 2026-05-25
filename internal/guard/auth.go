package guard

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)
		loggedIn := session.Get(
			"admin_logged_in",
		)

		if loggedIn != true {
			c.Redirect(
				http.StatusFound,
				"/login",
			)
			c.Abort()
			return
		}
		c.Next()
	}
}