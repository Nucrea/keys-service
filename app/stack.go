package app

import (
	"sync"
)

func NewStack[T any](initialCapacity int) *Stack[T] {
	return &Stack[T]{
		array: make([]T, 0, initialCapacity),
	}
}

type Stack[T any] struct {
	mutex sync.RWMutex
	array []T
}

func (s *Stack[T]) Push(val T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.array = append(s.array, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var retVal T
	if len(s.array) == 0 {
		return retVal, false
	}

	retVal = s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]

	return retVal, true
}

func (s *Stack[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.array)
}
