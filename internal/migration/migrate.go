package migration

import (
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func Run() {

	db.DB.AutoMigrate(
		&users.User{},
	)
}