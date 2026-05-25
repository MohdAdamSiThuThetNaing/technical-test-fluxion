package dashboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/users"
)

type TestResult struct {
	Action string `json:"Action"`
	Test   string `json:"Test"`
}

func Dashboard(c *gin.Context) {

	var totalUsers int64

	db.DB.Model(&users.User{}).
		Count(&totalUsers)

	file, _ := os.ReadFile(
		"test-results.json",
	)

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
		"dashboard.html",
		gin.H{
			"TotalUsers": totalUsers,
			"Results":    results,
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