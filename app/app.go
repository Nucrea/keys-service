package app

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App struct{}

func (a *App) Run(ctx context.Context) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Log().Msg("starting key service...")

	generator := NewRsaGenerator(2048)
	stack := NewStack[[]byte](5000)
	keysService := NewKeysService(&logger, generator, stack, 5000, 16)
	server := &Server{
		logger:      &logger,
		KeysService: keysService,
	}

	appCtx, cancelAppCtx := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSTOP, syscall.SIGINT)
	defer cancelAppCtx()

	workers := atomic.Int64{}
	workers.Store(2)

	go func() {
		err := keysService.Routine(appCtx)
		if err != nil {
			logger.Err(err).Msg("keys routine stopped with error")
		}
		workers.Add(-1)
	}()
	go func() {
		err := server.Run(appCtx, ":8080")
		if err != nil {
			logger.Err(err).Msg("server routine stopped with error")
		}
		workers.Add(-1)
	}()

	logger.Log().Msg("service started up")

	<-appCtx.Done()
	logger.Log().Msg("stopping service, waiting for goroutines to finish...")

	timestamp := time.Now().Add(time.Second)
	for {
		if time.Now().After(timestamp) {
			logger.Fatal().Msg("some routines not stopping, have to call os.Exit")
		}
		if workers.Load() <= 0 {
			break
		}
	}

	logger.Log().Msg("service stopped")
}
