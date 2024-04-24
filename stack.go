package main

import (
	"sync/atomic"
	"time"
)

func NewStack[T any](capacity int) *Stack[T] {
	spinlock := &atomic.Bool{}
	spinlock.Store(true)

	return &Stack[T]{
		spinlock: spinlock,
		array:    make([]T, 0, capacity),
	}
}

type Stack[T any] struct {
	spinlock *atomic.Bool
	array    []T
}

func (s *Stack[T]) Push(val T) {
	for !s.spinlock.Swap(false) {
		time.Sleep(time.Nanosecond)
	}
	defer s.spinlock.Swap(true)

	s.array = append(s.array, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	for !s.spinlock.Swap(false) {
		time.Sleep(time.Nanosecond)
	}
	defer s.spinlock.Swap(true)

	var retVal T
	if len(s.array) == 0 {
		return retVal, false
	}

	retVal = s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]

	return retVal, true
}

func (s *Stack[T]) Len() int {
	for !s.spinlock.Swap(false) {
		time.Sleep(time.Nanosecond)
	}
	defer s.spinlock.Swap(true)

	return len(s.array)
}
