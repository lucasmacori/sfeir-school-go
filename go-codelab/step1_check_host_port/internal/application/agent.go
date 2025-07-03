package application

import (
	"context"
	"time"
)

func InitAgent() AgentPort {
	return &agent{}
}

type agent struct {
}

func (a agent) Run(ctx context.Context) error {
	println("Chat with Claude (use 'ctrl-c') to quit")
	for {
		println("Claude : Wesh ma gueule, moi c'est Claude")
		time.Sleep(2 * time.Second)
	}
}
