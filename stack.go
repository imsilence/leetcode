package leetcode

import "strings"

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{make([]string, 0, 30)}
}

func (s *Stack) Push(e string) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (string, bool) {
	e, isEmpty := s.Top()
	if !isEmpty {
		s.elements = s.elements[:len(s.elements)-1]
	}
	return e, isEmpty
}

func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", true
	}
	return s.elements[len(s.elements)-1], false
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Clone() *Stack {
	ns := NewStack()
	ns.elements = append(ns.elements, s.elements...)
	return ns
}

func (s *Stack) String() string {
	return strings.Join(s.elements, "->")
}

func (s *Stack) Elements() string {
	return strings.Join(s.elements, "")
}

func (s *Stack) Clear() {
	s.elements = make([]string, 0, 30)
}
