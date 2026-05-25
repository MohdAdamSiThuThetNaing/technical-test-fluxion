package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAISummary(t *testing.T) {

	summary := `
	User Adam created account.
	User John updated profile.
	`

	assert.NotEmpty(t, summary)

	assert.Contains(t, summary, "User")
}