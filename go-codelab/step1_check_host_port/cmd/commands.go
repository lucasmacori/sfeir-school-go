package cmd

import (
	"go-agent/cmd/check_dns"
	llm_agent "go-agent/cmd/llm-agent"

	"github.com/spf13/cobra"
)

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(check_dns.New())
	rootCmd.AddCommand(llm_agent.New())
}
