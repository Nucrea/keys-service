package app

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

type confImpl struct{}

func (c *confImpl) GetThreadsCount() uint {
	return 16
}

func (c *confImpl) GetStackSize() uint {
	return 1000
}

func (c *confImpl) GetPort() uint16 {
	return 8081
}

func TestApp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := &confImpl{}
	logger := zerolog.New(os.Stdout)
	app := App{}

	go app.RunWithConfig(ctx, &logger, conf)
	time.Sleep(1 * time.Second)

	for i := 0; i < 200; i++ {
		resp, err := http.DefaultClient.Get("http://localhost:8081/health")
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)

		resp, err = http.DefaultClient.Get("http://localhost:8081/key")
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	}
}

func BenchmarkApp(t *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := &confImpl{}
	logger := zerolog.New(os.Stdout)
	app := App{}

	go app.RunWithConfig(ctx, &logger, conf)
	time.Sleep(1 * time.Second)

	for i := 0; i < t.N; i++ {
		resp, err := http.DefaultClient.Get("http://localhost:8080/key")
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	}
}
