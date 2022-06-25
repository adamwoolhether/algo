package algo

import (
	"sync"
)

// Stack uses an array as the underlying data structure.
type Stack struct {
	mu   sync.RWMutex
	data []any
}

func newStack() *Stack {
	return &Stack{
		mu:   sync.RWMutex{},
		data: []any{},
	}
}

func (s *Stack) push(v any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, v)
}

func (s *Stack) pop() any {
	if len(s.data) == 0 {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v
}

func (s *Stack) peak() any {
	if len(s.data) == 0 {
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[len(s.data)-1]
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Usecase: linter

type Linter struct {
	stack *Stack
}

func newLinter() *Linter {
	return &Linter{stack: newStack()}
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
			l.stack.push(char)
		case ')', ']', '}':
			poppedOpenBrace := l.stack.pop()

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
	if l.stack.peak() != nil {
		return false
	}

	return true
}
