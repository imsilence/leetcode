package leetcode

/*
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

type longestValidParenthesesElement struct {
	index int
	ch    rune
}

type longestValidParenthesesStack struct {
	elements []longestValidParenthesesElement
}

func newLongestValidParenthesesStack() *longestValidParenthesesStack {
	return &longestValidParenthesesStack{
		make([]longestValidParenthesesElement, 0, 30),
	}
}

func (s *longestValidParenthesesStack) Push(index int, ch rune) {
	s.elements = append(s.elements, longestValidParenthesesElement{index, ch})
}

func (s *longestValidParenthesesStack) Top() (int, rune, bool) {
	if len(s.elements) == 0 {
		return 0, 0, true
	}

	e := s.elements[len(s.elements)-1]
	return e.index, e.ch, false
}

func (s *longestValidParenthesesStack) Pop() (int, rune, bool) {
	if len(s.elements) == 0 {
		return 0, 0, true
	}

	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e.index, e.ch, false
}

func (s *longestValidParenthesesStack) Clear() {
	s.elements = make([]longestValidParenthesesElement, 0, 30)
}

func longestValidParentheses(s string) int {
	stack := newLongestValidParenthesesStack()
	maxLengths := make([]int, len(s))
	length := 0

	for i, ch := range s {
		if ch == '(' {
			stack.Push(i, ch)
		} else if _, top, _ := stack.Top(); top == '(' {
			stack.Pop()
			length += 2
		} else {
			length = 0
			stack.Clear()
		}
		maxLengths[i] = length
	}

	for {
		index, _, isEmpty := stack.Pop()
		if isEmpty {
			break
		}
		for i := index + 1; i < len(maxLengths); i++ {
			if maxLengths[i] == 0 {
				break
			}
			maxLengths[i] -= maxLengths[index]
		}
		maxLengths[index] = 0
	}

	maxLength := 0
	for _, length := range maxLengths {
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
