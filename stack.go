package leetcode

import "strings"

// Stack 字符串类型堆栈
type Stack struct {
	elements []string
}

// NewStack 创建字符串堆栈
func NewStack() *Stack {
	return &Stack{make([]string, 0, 30)}
}

// Push 入栈
func (s *Stack) Push(e string) {
	s.elements = append(s.elements, e)
}

// Pop 出栈
func (s *Stack) Pop() (string, bool) {
	e, isEmpty := s.Top()
	if !isEmpty {
		s.elements = s.elements[:len(s.elements)-1]
	}
	return e, isEmpty
}

// Top 获取栈顶元素
func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", true
	}
	return s.elements[len(s.elements)-1], false
}

// IsEmpty 检查是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Clone 复制堆栈元素
func (s *Stack) Clone() *Stack {
	ns := NewStack()
	ns.elements = append(ns.elements, s.elements...)
	return ns
}

func (s *Stack) String() string {
	return strings.Join(s.elements, "->")
}

// Elements 获取堆栈所有元素
func (s *Stack) Elements() string {
	return strings.Join(s.elements, "")
}

// Clear 清空堆栈
func (s *Stack) Clear() {
	s.elements = make([]string, 0, 30)
}
