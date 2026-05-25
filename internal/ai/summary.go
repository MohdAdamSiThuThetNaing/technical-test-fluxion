package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateSummary(logsText string) (string, error) {

	prompt := `
You are a backend log analyzer.

STRICT RULES:
- ONLY analyze provided logs
- DO NOT invent timestamps
- DO NOT invent users
- DO NOT invent events
- DO NOT hallucinate
- Return concise factual summaries only

Return format:

{
  "total_user_created": number,
  "total_user_updated": number,
  "total_user_deleted": number,
  "latest_admin_action": "EVENT_NAME"
}

If information is missing:
- use 0
- use "NOT_FOUND"

Application Logs:
` + logsText

	reqBody := OllamaRequest{
		Model:  "phi3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)

	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		os.Getenv("OLLAMA_URL")+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return "", fmt.Errorf(
			"ollama returned status: %d",
			resp.StatusCode,
		)
	}

	var result OllamaResponse
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(result.Response), nil
}