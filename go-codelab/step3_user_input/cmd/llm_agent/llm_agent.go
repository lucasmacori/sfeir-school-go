package llm_agent

import (
	"context"
	"fmt"
	"go-agent/internal/application"
	anthropicLocal "go-agent/internal/infrastructure/anthropic"

	"github.com/spf13/cobra"
)

var example = `# sfeir_cli llm-agent`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "llm-agent",
		Short:   "Run the LLM agent",
		Example: example,
		Run: func(cmd *cobra.Command, args []string) {
			agent := application.InitAgent(anthropicLocal.NewLLM())
			err := agent.Run(context.TODO())
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}

		},
	}

	return cmd
}
