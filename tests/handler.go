package tests

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TestResult struct {
	Action string `json:"Action"`
	Test   string `json:"Test"`
}

func ShowTests(c *gin.Context) {

	file, err := os.ReadFile(
		"test-results.json",
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

	lines := splitLines(string(file))

	testMap := make(map[string]TestResult)

	for _, line := range lines {

		var result TestResult

		err := json.Unmarshal(
			[]byte(line),
			&result,
		)

		if err == nil &&
			result.Test != "" &&
			(result.Action == "pass" ||
				result.Action == "fail") {

			testMap[result.Test] = result
		}
	}

	var results []TestResult

	for _, result := range testMap {

		results = append(
			results,
			result,
		)
	}

	c.HTML(
		http.StatusOK,
		"tests.html",
		gin.H{
			"Results": results,
		},
	)
}

func splitLines(s string) []string {

	var lines []string
	current := ""

	for _, ch := range s {

		if ch == '\n' {

			lines = append(lines, current)
			current = ""
			continue
		}

		current += string(ch)
	}

	return lines
}