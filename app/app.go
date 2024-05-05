package app

import (
	"context"
	"keys_service/args"
	"keys_service/config"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App struct{}

func (a *App) Run(ctx context.Context, cmdArgs []string) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Log().Msg("starting key service...")

	args, err := args.NewArgs(cmdArgs)
	if err != nil {
		logger.Fatal().Err(err).Msg("error parsing command line arguments")
	}

	conf, err := config.NewConfig(args.GetConfigFilePath())
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening config file")
	}

	a.RunWithConfig(ctx, &logger, conf)
}

func (a *App) RunWithConfig(ctx context.Context, logger *zerolog.Logger, conf config.Config) {
	generator := NewRsaGenerator(2048)
	stack := NewStack[[]byte](int(conf.GetStackSize()))
	keysService := NewKeysService(logger, generator, stack, int(conf.GetStackSize()), int(conf.GetThreadsCount()))
	server := &Server{
		logger:      logger,
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
		err := server.Run(appCtx, conf.GetPort())
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
