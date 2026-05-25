package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
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
You are a backend system log analyzer.

ONLY summarize the provided application logs.

DO NOT create fictional stories, games, characters, or assumptions.

Focus only on:
- user creation
- user update
- user deletion
- timestamps
- admin actions

Logs:
` + logsText

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
		return "", err
	}

	defer resp.Body.Close()
	var result OllamaResponse
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	return result.Response, nil
}