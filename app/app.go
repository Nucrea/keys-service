package app

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type App struct{}

func (a *App) Run(ctx context.Context) {
	logger := zerolog.New(os.Stdout)
	logger.Log().Msg("starting key service...")

	// newCtx, stopCtx := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM|syscall.SIGKILL|syscall.SIGSTOP|syscall.SIGINT)
	// defer stopCtx()

	generator := NewRsaGenerator(2048)
	stack := NewStack[[]byte](5000)
	keysService := NewKeysService(&logger, generator, stack, 5000, 16)

	go keysService.Routine(ctx)

	logger.Log().Msg("started listening on port :8080")

	server := Server{keysService}
	err := server.Run(ctx, ":8080")
	if err != nil {
		logger.Fatal().Err(err).Msg("server stopped with error")
	}
}
