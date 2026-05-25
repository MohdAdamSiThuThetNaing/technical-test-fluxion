package ai

import (
	"bytes"
	"encoding/json"
	"io"
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

func GenerateSummary(logs string) (string, error) {

	requestBody := OllamaRequest{
		Model: "phi3",
		Prompt: "Summarize these logs:\n\n" + logs,
		Stream: false,
	}

	jsonData, _ := json.Marshal(requestBody)

	ollamaURL := os.Getenv("OLLAMA_URL")

	resp, err := http.Post(
		ollamaURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result OllamaResponse

	json.Unmarshal(body, &result)

	return result.Response, nil
}