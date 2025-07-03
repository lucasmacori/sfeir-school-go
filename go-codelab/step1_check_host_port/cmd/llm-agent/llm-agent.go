package llm_agent

import (
	"context"
	"go-agent/internal/application"

	"github.com/spf13/cobra"
)

var example = `# llm-agent`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "llm-agent",
		Short:   "Run the LLM agent",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {
			agent := application.InitAgent()
			if err := agent.Run(context.TODO()); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
