package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/ai"
)

func TestGenerateSummary(t *testing.T) {

	os.Setenv(
		"OLLAMA_URL",
		"http://localhost:11434",
	)

	result, err := ai.GenerateSummary(
		"USER_CREATED user@gmail.com",
	)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}