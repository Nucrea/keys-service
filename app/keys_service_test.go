package app

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestKeysServiceGetKey(t *testing.T) {
	logger := zerolog.New(os.Stdout)
	stack := NewStack[[]byte](1)
	generator := NewRsaGenerator(2048)
	keysService := NewKeysService(&logger, generator, stack, 1, 1)

	testBytes := []byte("test bytes")
	stack.Push(testBytes)

	result, ok := keysService.GetKey()
	require.True(t, ok)
	require.Equal(t, testBytes, result)

	result, ok = keysService.GetKey()
	require.False(t, ok)
	require.Nil(t, result)
}

func TestKeysServiceRoutine(t *testing.T) {
	keysCount := 100

	logger := zerolog.New(os.Stdout)
	stack := NewStack[[]byte](keysCount)
	generator := NewRsaGenerator(2048)
	keysService := NewKeysService(&logger, generator, stack, keysCount, 16)

	go keysService.Routine(context.Background())

	for stack.Len() < keysCount {
		time.Sleep(time.Millisecond)
	}
}

func BenchmarkKeysService(t *testing.B) {
	logger := zerolog.New(os.Stdout)
	stack := NewStack[[]byte](t.N)
	generator := NewRsaGenerator(2048)
	keysService := NewKeysService(&logger, generator, stack, t.N, 16)

	go keysService.Routine(context.Background())

	for stack.Len() < t.N {
		time.Sleep(time.Millisecond)
	}
}
