package application

import (
	"context"
	"go-agent/internal/domain"
)

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
}

type AgentPort interface {
	Run(ctx context.Context) error
}
