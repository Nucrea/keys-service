package app

import (
	"bytes"
	"context"
	"net"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type IKeysService interface {
	GetKey() ([]byte, bool)
}

type Server struct {
	KeysService IKeysService
}

func (s *Server) Run(ctx context.Context, addr string) error {
	listenCfg := net.ListenConfig{}
	listener, err := listenCfg.Listen(ctx, "tcp4", addr)
	if err != nil {
		return err
	}

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
