package ai

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"
)

func TestAI(c *gin.Context) {

	collection := db.MongoClient.
		Database("fluxion_logs").
		Collection("logs")

	cursor, err := collection.Find(
		context.Background(),
		map[string]interface{}{},
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	defer cursor.Close(
		context.Background(),
	)

	var logsText string
	for cursor.Next(
		context.Background(),
	) {

		var logData logs.Log
		cursor.Decode(&logData)
		logsText +=
			"Event: " + logData.Event +
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

	input := c.PostForm("input")
	if input == "" {
		c.HTML(
			http.StatusBadRequest,
			"ai-suggestion.html",
			gin.H{
				"Error": "prompt is required",
			},
		)
		return
	}

	result, err := GenerateUserSuggestion(input)
	if err != nil {
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
		200,
		"ai-suggestion.html",
		nil,
	)
}