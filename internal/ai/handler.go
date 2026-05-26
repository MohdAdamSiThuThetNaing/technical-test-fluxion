package ai

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/audit"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
)

type AISuggestionReason string

const (

	EMPTY_PROMPT AISuggestionReason = "EMPTY_PROMPT"
)

func TestAI(c *gin.Context) {

	collection := db.MongoClient.Database("fluxion_logs").Collection("logs")
	cursor, err := collection.Find(context.Background(), map[string]interface{}{})

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	defer cursor.Close(context.Background())
	var logsText string

	for cursor.Next(context.Background()) {

		var logData logs.Log
		cursor.Decode(&logData)
		logsText += "Event: " + logData.Event +
			", Email: " + logData.UserEmail +
			", Name: " + logData.UpdatedName +
			"\n"
	}

	result, err := GenerateSummary(logsText)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"result": result,
		},
	)
}

func SuggestUser(c *gin.Context) {

	session := sessions.Default(c)
	
	userID := toString(session.Get("user_id"))
	userEmail := toString(session.Get("user_email"))
	userName := toString(session.Get("user_name"))
	input := c.PostForm("input")

	if input == "" {

		audit.PublishAuditLog(
			c,
			audit.AI_SUGGESTION_FAILED,
			userID,
			userEmail,
			userName,
			string(EMPTY_PROMPT),
		)

		c.HTML(
			http.StatusBadRequest,
			"ai-suggestion.html",
			gin.H{
				"Error": "prompt is required",
			},
		)

		return
	}

	audit.PublishAuditLog(
		c,
		audit.AI_SUGGESTION_STARTED,
		userID,
		userEmail,
		userName,
		input,
	)

	result, err := GenerateUserSuggestion(input)

	if err != nil {

		audit.PublishAuditLog(
			c,
			audit.AI_SUGGESTION_FAILED,
			userID,
			userEmail,
			userName,
			err.Error(),
		)

		c.HTML(
			http.StatusInternalServerError,
			"ai-suggestion.html",
			gin.H{
				"Error": err.Error(),
				"Input": input,
			},
		)

		return
	}

	audit.PublishAuditLog(
		c,
		audit.AI_SUGGESTION_SUCCESS,
		userID,
		userEmail,
		userName,
		result.Role,
	)

	c.HTML(
		http.StatusOK,
		"ai-suggestion.html",
		gin.H{
			"Result": result,
			"Input":  input,
		},
	)
}

func ShowUserSuggestion(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"ai-suggestion.html",
		nil,
	)
}

func toString(value interface{}) string {

	if value == nil {
		return ""
	}

	str, ok := value.(string)
	if !ok {
		return ""
	}
	return str
}