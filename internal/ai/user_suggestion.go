package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type UserSuggestion struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func GenerateUserSuggestion(input string) (*UserSuggestion, error) {

	prompt := `
Generate a professional user profile.

Return ONLY valid JSON.

Format:
{
  "name": "",
  "email": "",
  "role": ""
}

Input:
` + input

	reqBody := OllamaRequest{
		Model:  "phi3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		os.Getenv("OLLAMA_URL")+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf(
			"ollama error: %d",
			resp.StatusCode,
		)
	}

	var ollamaResp OllamaResponse

	json.NewDecoder(resp.Body).Decode(&ollamaResp)

	cleanJSON := strings.TrimSpace(ollamaResp.Response)
	cleanJSON = strings.ReplaceAll(cleanJSON, "```json", "")
	cleanJSON = strings.ReplaceAll(cleanJSON, "```", "")

	var user UserSuggestion

	err = json.Unmarshal(
		[]byte(cleanJSON),
		&user,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}