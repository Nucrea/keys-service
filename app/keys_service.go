package app

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
)

func NewKeysService(
	logger *zerolog.Logger,
	rsaGenerator IRsaGenerator,
	stack *Stack[[]byte],
	capacity, maxThreads int,
) *KeysService {
	return &KeysService{
		logger:       logger,
		rsaGenerator: rsaGenerator,
		stack:        stack,
		capacity:     capacity,
		maxThreads:   maxThreads,
	}
}

type KeysService struct {
	logger       *zerolog.Logger
	rsaGenerator IRsaGenerator
	stack        *Stack[[]byte]
	capacity     int
	maxThreads   int
}

func (k *KeysService) GetKey() ([]byte, bool) {
	return k.stack.Pop()
}

func (k *KeysService) Routine(ctx context.Context) {
	logTicker := time.NewTicker(10 * time.Second)
	errorChan := make(chan error, k.maxThreads)
	atmThreads := atomic.Int64{}
	errorsCount := 0

	defer func() {
		for atmThreads.Load() > 0 {
			k.idle()
		}
		close(errorChan)
		logTicker.Stop()
	}()

	k.logger.Log().Msgf("keys routine started, threads=%d, capacity=%d", k.maxThreads, k.capacity)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case err := <-errorChan:
			errorsCount++
			k.logger.Err(err).Msg("an error occured on generating rsa key (KeysService routine)")
		default:
		}

		threads := int(atmThreads.Load())
		stackLen := k.stack.Len()

		select {
		case <-logTicker.C:
			k.logger.Log().Msgf("working threads: %d, stack size: %d, errors count: %d", threads, stackLen, errorsCount)
		default:
		}

		if stackLen+threads >= k.capacity || threads >= k.maxThreads {
			k.idle()
			continue
		}

		atmThreads.Add(1)
		go k.keyJob(&atmThreads, &errorChan)
	}
}

func (k *KeysService) idle() {
	time.Sleep(time.Millisecond)
}

func (k *KeysService) keyJob(atmThreads *atomic.Int64, errorChan *chan error) {
	defer func() {
		atmThreads.Add(-1)
	}()

	key, err := k.rsaGenerator.NewKey()
	if err != nil {
		*errorChan <- fmt.Errorf("error generating rsa key: %w", err)
		return
	}

	k.stack.Push(key)
}
