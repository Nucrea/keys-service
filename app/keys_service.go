package app

import (
	"context"
	"sync/atomic"
	"time"
)

func NewKeysService(
	rsaGenerator IRsaGenerator,
	stack *Stack[[]byte],
	capacity, maxThreads int,
) *KeysService {
	return &KeysService{
		rsaGenerator: rsaGenerator,
		stack:        stack,
		capacity:     capacity,
		maxThreads:   maxThreads,
	}
}

type KeysService struct {
	rsaGenerator IRsaGenerator
	stack        *Stack[[]byte]
	capacity     int
	maxThreads   int
}

func (k *KeysService) GetKey() ([]byte, bool) {
	return k.stack.Pop()
}

func (k *KeysService) Routine(ctx context.Context) {
	atmThreads := atomic.Int64{}
	defer func() {
		for atmThreads.Load() > 0 {
			//NOP
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		threads := int(atmThreads.Load())
		if k.stack.Len()+threads >= k.capacity || threads >= k.maxThreads {
			time.Sleep(time.Millisecond)
			continue
		}

		atmThreads.Add(1)
		go k.keyJob(&atmThreads)
	}
}

func (k *KeysService) keyJob(atmThreads *atomic.Int64) {
	defer func() {
		atmThreads.Add(-1)
	}()

	key, _ := k.rsaGenerator.NewKey()
	// if err != nil {
	// 	return fmt.Errorf("error generating rsa key: %w", err)
	// }

	k.stack.Push(key)
	// return nil
}
