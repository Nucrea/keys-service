package app

import (
	"context"
	"fmt"
)

type App struct{}

func (a *App) Run(ctx context.Context) {
	const stackLen = 100

	stack := NewStack[[]byte](stackLen)
	rsaGenerator := RSAGenerator{}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if stack.Len() >= stackLen {
			continue
		}

		key, err := rsaGenerator.GenerateRSAKey()
		if err != nil {
			fmt.Printf("error generating rsa key: %s\n", err.Error())
		}

		stack.Push(key)
		fmt.Printf("key generated (stack size: %d)\n", stack.Len())
	}
}
