package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func TestUserModel(t *testing.T) {

	user := users.User{
		Name:  "Adam",
		Email: "adam@gmail.com",
	}

	assert.Equal(
		t,
		"Adam",
		user.Name,
	)

	assert.Equal(
		t,
		"adam@gmail.com",
		user.Email,
	)
}