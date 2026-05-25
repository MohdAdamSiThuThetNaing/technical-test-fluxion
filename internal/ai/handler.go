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

	type Request struct {
		Prompt string `json:"prompt"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if req.Prompt == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "prompt is required",
			},
		)
		return
	}

	result, err := GenerateUserSuggestion(req.Prompt)
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
			"suggestion": result,
		},
	)
}