package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	LLMAPIBaseURL = os.Getenv("LLM_API_BASE_URL")
	LLMAPIKey     = os.Getenv("LLM_API_KEY")
	LLMChatModel  = os.Getenv("LLM_CHAT_MODEL")
	LLMDebug      = os.Getenv("LLM_DEBUG")
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Stop        []string  `json:"stop"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type Choice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func chat(messages []Message) (string, error) {
	url := fmt.Sprintf("%s/chat/completions", LLMAPIBaseURL)
	authHeader := ""
	if LLMAPIKey != "" {
		authHeader = fmt.Sprintf("Bearer %s", LLMAPIKey)
	}
	requestBody := ChatRequest{
		Messages:    messages,
		Model:       LLMChatModel,
		Stop:        []string{"<|im_end|>", "<|end|>", "<|eot_id|>"},
		MaxTokens:   200,
		Temperature: 0,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	var data struct {
		Choices []Choice `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	answer := data.Choices[0].Message.Content
	return answer, nil
}

const SystemPrompt = "Answer the question politely and concisely."

func main() {
	fmt.Printf("Using LLM at %s.\n", LLMAPIBaseURL)
	fmt.Println("Press Ctrl+D to exit.")
	fmt.Println()

	messages := []Message{{Role: "system", Content: SystemPrompt}}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")
		scanner.Scan()
		question := scanner.Text()

		if question == "" {
			break
		}

		messages = append(messages, Message{Role: "user", Content: question})
		start := time.Now()
		answer, err := chat(messages)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		messages = append(messages, Message{Role: "assistant", Content: answer})
		fmt.Println(answer)
		elapsed := time.Since(start)
		if LLMDebug != "" {
			fmt.Printf("[%d ms]\n", elapsed.Milliseconds())
		}
		fmt.Println()
	}
}
