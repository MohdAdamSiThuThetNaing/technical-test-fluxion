package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/ai"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
)

type EnvironmentKey string

const (

	OLLAMA_URL_KEY EnvironmentKey = "OLLAMA_URL"
)

type EnvironmentValue string

const (

	LOCAL_OLLAMA_URL EnvironmentValue = "http://localhost:11434"
)

type SummaryTestData string

const (

	TEST_EMAIL SummaryTestData = "user@gmail.com"
)

func TestGenerateSummary(t *testing.T) {

	os.Setenv(
		string(OLLAMA_URL_KEY),
		string(LOCAL_OLLAMA_URL),
	)

	logText :=
		string(audit.USER_CREATED) +
			" " +
			string(TEST_EMAIL)

	result, err := ai.GenerateSummary(logText)

	assert.Nil(
		t,
		err,
	)

	assert.NotEmpty(
		t,
		result,
	)
}