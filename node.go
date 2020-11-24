package leetcode

import (
	"fmt"
	"strings"
)

// ListNode Node
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Equals(r *ListNode) bool {
	head := n
	for head != nil && r != nil {
		if head.Val != r.Val {
			return false
		}
		head = head.Next
		r = r.Next
	}

	return head == nil && r == nil
}

func (n *ListNode) String() string {
	head := n
	var builder strings.Builder
	for head != nil {
		builder.WriteString(fmt.Sprintf("%d", head.Val))
		builder.WriteString("->")
		head = head.Next
	}

	builder.WriteString("nil")
	return builder.String()
}
