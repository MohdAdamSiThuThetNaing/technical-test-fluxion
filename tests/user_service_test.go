package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

func TestCreateUserDTO(t *testing.T) {

	input := users.CreateUserDTO{
		Name:     "Adam",
		Email:    "adam@example.com",
		Password: "password123",
	}

	assert.Equal(t, "Adam", input.Name)

	assert.Equal(
		t,
		"adam@example.com",
		input.Email,
	)

	assert.NotEmpty(
		t,
		input.Password,
	)
}