package algo

import (
	"sync"
)

// Stack uses an array as the underlying data structure.
type Stack[T any] struct {
	mu   sync.RWMutex
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		mu:   sync.RWMutex{},
		data: []T{},
	}
}

func (s *Stack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isEmpty() {
		return s.nilType()
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v
}

func (s *Stack[T]) Peak() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.isEmpty() {
		return s.nilType()
	}

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

func (s *Stack[T]) isEmpty() bool {
	// no locking, as this is internal and should be done by caller.
	return len(s.data) == 0
}

func (s *Stack[T]) nilType() T {
	var t T

	return t
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Usecase: linter

type Linter struct {
	stack *Stack[any]
}

func newLinter() *Linter {
	return &Linter{stack: NewStack[any]()}
}
func (l *Linter) lint(text string) bool {
	brace := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	// Loop through each char.
	for _, char := range text {
		switch char {
		case '(', '[', '{':
			l.stack.Push(char)
		case ')', ']', '}':
			poppedOpenBrace := l.stack.Pop()

			// If the stack was empty.
			if poppedOpenBrace == nil {
				return false
			}

			// If the popped brace doesn't correspond to the closing brace.
			if check := brace[char]; check != poppedOpenBrace {
				return false
			}
		}
	}

	// There are remaining chars in the stack, it lacks a closing brace.
	if l.stack.Peak() != nil {
		return false
	}

	return true
}
