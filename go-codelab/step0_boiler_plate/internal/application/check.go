package application

import (
	"fmt"
	"go-agent/internal/domain"
	"net"
)

func InitCheck() CheckerPort {
	return &Checker{}

}

type Checker struct{}

func (c Checker) ResolveDNS(opts domain.CheckHostOptions) error {
	ips, err := net.LookupHost(opts.Host)
	if err != nil {
		return fmt.Errorf("Lookup host: %w", err)
	}
	fmt.Printf("DNS resolution for host %s\n", opts.Host)
	for _, ip := range ips {
		fmt.Printf("- %v", ip)
	}
	return nil
}
