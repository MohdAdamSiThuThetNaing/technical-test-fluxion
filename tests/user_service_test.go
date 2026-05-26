package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func TestUserModel(t *testing.T) {

	expectedName := "AdamNyi"

	expectedEmail := "adamnyi@gmail.com"

	currentTime := time.Now()

	user := users.User{
		ID:        uuid.New(),
		Name:      expectedName,
		Email:     expectedEmail,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	assert.NotEmpty(t, user.ID)

	assert.Equal(t, expectedName, user.Name)
	assert.Equal(t, expectedEmail, user.Email)
	assert.WithinDuration(t, currentTime, user.CreatedAt, time.Second)
	assert.WithinDuration(t, currentTime, user.UpdatedAt, time.Second)
}