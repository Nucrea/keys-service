package app

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	app := App{}
	go app.Run(context.Background())
	time.Sleep(3 * time.Second)

	for i := 0; i < 200; i++ {
		resp, err := http.DefaultClient.Get("http://localhost:8080/health")
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)

		resp, err = http.DefaultClient.Get("http://localhost:8080/key")
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	}
}
