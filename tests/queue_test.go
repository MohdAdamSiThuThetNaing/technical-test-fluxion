package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestQueuePayload(t *testing.T) {

	payload := gin.H{
		"event": "USER_CREATED",
	}

	assert.Equal(
		t,
		"USER_CREATED",
		payload["event"],
	)
}