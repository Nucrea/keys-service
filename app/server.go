package app

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type IKeysService interface {
	GetKey() ([]byte, bool)
}

type Server struct {
	logger      *zerolog.Logger
	KeysService IKeysService
}

func (s *Server) Run(ctx context.Context, port uint16) error {
	addr := fmt.Sprintf(":%d", port)

	listenCfg := net.ListenConfig{}
	listener, err := listenCfg.Listen(ctx, "tcp4", addr)
	if err != nil {
		return err
	}
	s.logger.Log().Msgf("server listener initialized on %s", addr)

	finishedChan := make(chan struct{})
	go func() {
		<-ctx.Done()
		listener.Close()
		close(finishedChan)
	}()

	defer func() {
		s.logger.Log().Msgf("stopping server, waiting for listener to close...")

		timeoutTicker := time.NewTicker(1 * time.Second)
		defer timeoutTicker.Stop()

		select {
		case <-timeoutTicker.C:
		case <-finishedChan:
		}
		s.logger.Log().Msgf("server stopped")
	}()

	return fasthttp.Serve(listener, func(ctx *fasthttp.RequestCtx) {
		switch {
		case bytes.Equal([]byte("/health"), ctx.Path()):
			s.getHealthHandler(ctx)

		case bytes.Equal([]byte("/key"), ctx.Path()):
			s.getKeyHandler(ctx)

		default:
			ctx.SetStatusCode(http.StatusNotFound)
		}
	})
}

func (s *Server) getHealthHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
}

func (s *Server) getKeyHandler(ctx *fasthttp.RequestCtx) {
	for i := 0; i < 1200; i++ {
		key, ok := s.KeysService.GetKey()
		if ok {
			ctx.Write(key)
			ctx.SetStatusCode(http.StatusOK)
			return
		}
		time.Sleep(50 * time.Millisecond)
	}

	ctx.SetStatusCode(http.StatusTooManyRequests)
}
