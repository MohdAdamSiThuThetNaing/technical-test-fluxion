package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
)

type QueueTestEvent string

const (

	TEST_QUEUE_USER_CREATED QueueTestEvent = "USER_CREATED"
)

func TestQueuePayload(t *testing.T) {

	expectedEvent := audit.USER_CREATED

	payload := gin.H{
		"event": string(expectedEvent),
	}

	assert.Equal(
		t,
		string(TEST_QUEUE_USER_CREATED),
		payload["event"],
	)
}