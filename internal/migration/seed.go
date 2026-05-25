package migration

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func SeedAdmin() {

	var existingUser users.User
	err := db.DB.
		Where("email = ?", "admin@fluxion.com").
		First(&existingUser).Error

	if err == nil {

		log.Println("Admin already seeded")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Fatal(err)
	}

	admin := users.User{
		ID:       uuid.New(),
		Name:     "System Admin",
		Email:    "admin@fluxion.com",
		Password: string(hashedPassword),
	}

	db.DB.Create(&admin)
	log.Println("Default admin seeded")
}