package main

import (
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/ai"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/auth"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/dashboard"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/migration"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/tests"
)

func main() {

	db.ConnectPostgres()
	db.ConnectMongo()

	migration.Run()
	migration.SeedAdmin()

	r := gin.Default()

	store := cookie.NewStore(
		[]byte("super-secret-session"),
	)

	r.Use(
		sessions.Sessions(
			"fluxion_session",
			store,
		),
	)

	// Load root-level templates
	tmpl := template.Must(
		template.New("").ParseGlob(
			"templates/*.html",
		),
	)

	// Load nested templates
	template.Must(
		tmpl.ParseGlob(
			"templates/*/*.html",
		),
	)

	r.SetHTMLTemplate(tmpl)

	auth.RegisterRoutes(r)
	dashboard.RegisterRoutes(r)
	users.RegisterRoutes(r)
	logs.RegisterRoutes(r)
	ai.RegisterRoutes(r)
	tests.RegisterRoutes(r)

	r.Run(":8080")
}