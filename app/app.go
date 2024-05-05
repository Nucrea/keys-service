package app

import (
	"context"
)

type App struct{}

func (a *App) Run(ctx context.Context) {
	generator := NewRsaGenerator(2048)
	stack := NewStack[[]byte](5000)
	keysService := NewKeysService(generator, stack, 5000, 16)

	go keysService.Routine(ctx)

	server := Server{keysService}
	err := server.Run(ctx, ":8080")
	if err != nil {
		panic(err)
	}
}
