package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
)

func TestLogsModel(t *testing.T) {

	logData := logs.Log{
		Event: "USER_CREATED",
	}

	assert.Equal(
		t,
		"USER_CREATED",
		logData.Event,
	)
}