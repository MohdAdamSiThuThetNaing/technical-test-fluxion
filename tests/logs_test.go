package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
)

type LogsTestEvent string

const (

	TEST_LOGS_USER_CREATED LogsTestEvent = "USER_CREATED"
)

func TestLogsModel(t *testing.T) {

	expectedEvent := audit.USER_CREATED

	logData := logs.Log{
		Event: string(expectedEvent),
	}

	assert.Equal(
		t,
		string(TEST_LOGS_USER_CREATED),
		logData.Event,
	)
}