package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockKeysService struct {
	len int
}

func (k *MockKeysService) GetKey() ([]byte, bool) {
	if k.len == 0 {
		return nil, false
	}

	k.len--
	return []byte("key"), true
}

func TestServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := uint16(8080)

	errChan := make(chan error, 1)

	logger := zerolog.New(os.Stdout)
	server := Server{&logger, &MockKeysService{1}}
	go func() {
		errChan <- server.Run(ctx, port)
	}()

	select {
	case <-ctx.Done():
		return
	case err := <-errChan:
		t.Fatal(err)
	case <-time.NewTimer(time.Second).C:
	}

	addr := fmt.Sprintf("http://localhost:%d", port)

	t.Run("health route returns 200", func(t *testing.T) {
		resp, err := http.DefaultClient.Get(addr + "/health")
		assert.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("key route returns 200", func(t *testing.T) {
		resp, err := http.DefaultClient.Get(addr + "/key")
		assert.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("key route returns 502", func(t *testing.T) {
		resp, err := http.DefaultClient.Get(addr + "/key")
		assert.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 502, resp.StatusCode)
	})
}
