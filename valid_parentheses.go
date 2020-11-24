package leetcode

import "fmt"

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/

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

func isValid(s string) bool {
	stack := NewStack()
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack.Push(fmt.Sprintf("%c", ch))
		case ')':
			if e, _ := stack.Pop(); e != "(" {
				return false
			}
		case ']':
			if e, _ := stack.Pop(); e != "[" {
				return false
			}
		case '}':
			if e, _ := stack.Pop(); e != "{" {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
