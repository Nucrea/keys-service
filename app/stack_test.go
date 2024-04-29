package app

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	stack := NewStack[int](10)

	require.Equal(t, 0, stack.Len())

	valToPush := 111
	stack.Push(valToPush)

	require.Equal(t, 1, stack.Len())

	val, ok := stack.Pop()
	require.True(t, ok)
	require.Equal(t, valToPush, val)

	val, ok = stack.Pop()
	require.False(t, ok)
	require.Equal(t, 0, val)

	require.Equal(t, 0, stack.Len())
}

func TestStackConcurrent(t *testing.T) {
	stack := NewStack[int](10)

	maxCount := 10000
	counter := atomic.Int64{}
	getNextCount := func() int {
		return int(counter.Add(1) - 1)
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				nextVal := getNextCount()
				if nextVal >= maxCount {
					return
				}

				stack.Push(nextVal)
			}
		}()
	}
	wg.Wait()

	require.Equal(t, maxCount, stack.Len())

	valuesMap := map[int]bool{}
	for i := 0; i < maxCount; i++ {
		valuesMap[i] = true
	}

	for i := maxCount - 1; i >= 0; i-- {
		val, ok := stack.Pop()
		require.True(t, ok)
		require.True(t, valuesMap[val])

		valuesMap[val] = false
	}

	val, ok := stack.Pop()
	require.False(t, ok)
	require.Equal(t, 0, val)
}
