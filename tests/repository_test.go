package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
)

type LogTestEvent string

const (

	TEST_USER_CREATED LogTestEvent = "USER_CREATED"
)

func TestRepositoryLogModel(t *testing.T) {

	expectedEvent := audit.USER_CREATED

	logData := logs.Log{
		Event: string(expectedEvent),
	}

	assert.Equal(
		t,
		string(TEST_USER_CREATED),
		logData.Event,
	)
}