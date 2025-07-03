package application

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

func InitAgent(port LLMPort) AgentPort {
	return &Agent{
		llmPort: port,
	}
}

type Agent struct {
	llmPort LLMPort
}

func (a *Agent) Run(ctx context.Context) error {
	conv := []anthropic.MessageParam{}
	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}

		msg := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
		conv = append(conv, msg)

		inf, err := a.llmPort.RunInference(ctx, conv)
		if err != nil {
			return err
		}
		conv = append(conv, inf.ToParam())

		for _, content := range inf.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[Claude\u001b[0m: %v\n", content.Text)
			}
		}
	}

	return nil
}

func (a *Agent) getUserMessage() (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", false
	}
	return scanner.Text(), true
}
