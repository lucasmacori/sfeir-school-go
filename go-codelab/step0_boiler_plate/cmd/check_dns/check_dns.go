package check_dns

import (
	"fmt"
	"go-agent/cmd/cliflags"
	"go-agent/internal/application"
	"go-agent/internal/domain"

	"github.com/spf13/cobra"
)

var example = `# sfeir_cli check-dns google.com`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check-dns",
		Short:   "Resolve DNS for a host",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {

			flags, err := parseFlags(cmd)
			if err != nil {
				return err
			}

			check := application.InitCheck()
			if err := check.ResolveDNS(flags); err != nil {
				fmt.Println("Error %e", err)
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.String(cliflags.Host, "", "Focus the host")
	cmd.MarkFlagRequired(cliflags.Host)
	return cmd
}

func parseFlags(cmd *cobra.Command) (domain.CheckHostOptions, error) {
	flags := cmd.Flags()

	host, err := flags.GetString(cliflags.Host)
	if err != nil {
		return domain.CheckHostOptions{}, err
	}
	return domain.CheckHostOptions{Host: host}, nil
}
