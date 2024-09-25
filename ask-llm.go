package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	LLMAPIBaseURL = os.Getenv("LLM_API_BASE_URL")
	LLMAPIKey     = os.Getenv("LLM_API_KEY")
	LLMChatModel  = os.Getenv("LLM_CHAT_MODEL")
	LLMStreaming  = os.Getenv("LLM_STREAMING") != "no"
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
	Stream      bool      `json:"stream"`
}

type Choice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func chat(messages []Message, handler func(string)) (string, error) {
	url := fmt.Sprintf("%s/chat/completions", LLMAPIBaseURL)
	authHeader := ""
	if LLMAPIKey != "" {
		authHeader = fmt.Sprintf("Bearer %s", LLMAPIKey)
	}
	stream := LLMStreaming && handler != nil
	requestBody := ChatRequest{
		Messages:    messages,
		Model:       LLMChatModel,
		Stop:        []string{"<|im_end|>", "<|end|>", "<|eot_id|>"},
		MaxTokens:   200,
		Temperature: 0,
		Stream:      stream,
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

	if !stream {
		var data struct {
			Choices []Choice `json:"choices"`
		}
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return "", err
		}
		answer := data.Choices[0].Message.Content
		if handler != nil {
			handler(answer)
		}
		return answer, nil
	} else {
		answer := ""
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data: ") {
				payload := line[6:]
				var data struct {
					Choices []struct {
						Delta struct {
							Content string `json:"content"`
						} `json:"delta"`
					} `json:"choices"`
				}
				err := json.Unmarshal([]byte(payload), &data)
				if err != nil {
					return "", err
				}
				partial := data.Choices[0].Delta.Content
				answer += partial
				if handler != nil {
					handler(partial)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return answer, nil
	}
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
		handler := func(partial string) {
			fmt.Print(partial)
		}
		start := time.Now()
		answer, err := chat(messages, handler)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		messages = append(messages, Message{Role: "assistant", Content: answer})
		fmt.Println()
		elapsed := time.Since(start)
		if LLMDebug != "" {
			fmt.Printf("[%d ms]\n", elapsed.Milliseconds())
		}
		fmt.Println()
	}
}
