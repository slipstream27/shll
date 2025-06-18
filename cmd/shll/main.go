// main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: shll <natural language task>")
		os.Exit(1)
	}

	query := strings.Join(os.Args[1:], " ")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Set OPENAI_API_KEY env variable")
		os.Exit(1)
	}

	client := openai.NewClient(apiKey)
	maxResults := 3

	for {
		var resp openai.ChatCompletionResponse
		err := withLoadingIndicator("Waiting for command suggestions", func() error {
			var apiErr error
			resp, apiErr = client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT4oMini,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    "system",
							Content: fmt.Sprintf("You are a CLI assistant. Return %d shell command suggestions based on user input. No explanations.", maxResults),
						},
						{
							Role:    "user",
							Content: query,
						},
					},
				},
			)
			return apiErr
		})
		if err != nil {
			fmt.Println("Error calling OpenAI API:", err)
			os.Exit(1)
		}

		content := resp.Choices[0].Message.Content
		commands := parseSuggestions(content)

		var options []string
		options = append(options, commands...)

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("\nAvailable commands:")
			for i, opt := range options {
				num := color.New(color.FgCyan, color.Bold).SprintfFunc()
				cmd := color.New(color.FgGreen).SprintfFunc()
				fmt.Printf("%s %s\n", num("%d:", i+1), cmd(opt))
			}
			fmt.Println("m: More results")
			fmt.Print("\nSelect a command by number to copy to clipboard (add '?' to explain, e.g. 1?): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "q" || input == "Q" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}

			if input == "m" || input == "M" {
				maxResults += 3
				break // break inner loop to fetch more results
			}

			if strings.HasSuffix(input, "?") {
				idxStr := strings.TrimSuffix(input, "?")
				idx, err := strconv.Atoi(idxStr)
				if err != nil || idx < 1 || idx > len(options) {
					fmt.Println("Invalid selection.")
					continue
				}
				cmd := options[idx-1]
				explanation, err := explainCommand(client, cmd)
				if err != nil {
					fmt.Println("Error getting explanation:", err)
					continue
				}
				fmt.Printf("\nExplanation for: %s\n%s\n", cmd, explanation)
				fmt.Print("\nCopy this command to clipboard? No returns to previous options (y/n): ")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(strings.ToLower(choice))
				if choice == "y" {
					err = clipboard.WriteAll(cmd)
					if err != nil {
						fmt.Println("Failed to copy command to clipboard:", err)
						return
					}
					fmt.Println("\nThe command has been copied to your clipboard.")
					return
				}
				continue
			}

			idx, err := strconv.Atoi(input)
			if err != nil || idx < 1 || idx > len(options) {
				fmt.Println("Invalid selection.")
				continue
			}
			cmd := options[idx-1]
			err = clipboard.WriteAll(cmd)
			if err != nil {
				fmt.Println("Failed to copy command to clipboard:", err)
				return
			}
			fmt.Println("\nThe command has been copied to your clipboard.")
			return
		}
	}
}

func parseSuggestions(content string) []string {
	lines := strings.Split(content, "\n")
	var commands []string
	for _, line := range lines {
		line = strings.TrimPrefix(line, "1. `")
		line = strings.TrimPrefix(line, "2. `")
		line = strings.TrimPrefix(line, "3. `")
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "`")
		if line != "" {
			commands = append(commands, line)
		}
	}
	return commands
}

// Loading indicator utility
func withLoadingIndicator(message string, fn func() error) error {
	done := make(chan struct{})
	var err error
	go func() {
		fmt.Printf("%s", message)
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf(".")
			}
		}
	}()
	err = fn()
	close(done)
	return err
}

func explainCommand(client *openai.Client, cmd string) (string, error) {
	var explanation string
	err := withLoadingIndicator("Waiting for explanation", func() error {
		resp, apiErr := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oMini,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "system",
						Content: "You are a CLI assistant. Briefly and concisely explain the following shell command in a few sentences, focusing only on what it does. Explain what each executable is.",
					},
					{
						Role:    "user",
						Content: cmd,
					},
				},
			},
		)
		if apiErr != nil {
			return apiErr
		}
		explanation = resp.Choices[0].Message.Content
		return nil
	})
	if err != nil {
		return "", err
	}
	return explanation, nil
}
