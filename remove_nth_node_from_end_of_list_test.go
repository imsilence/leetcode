package leetcode

import "testing"

type RemoveNthNodeFromEndOfListCase struct {
	name   string
	list   *ListNode
	target int
	result *ListNode
}

func (c *RemoveNthNodeFromEndOfListCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createRemoveNthNodeFromEndOfListTestCase(t *testing.T, c *RemoveNthNodeFromEndOfListCase) {
	t.Helper()
	result := removeNthFromEnd(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []*RemoveNthNodeFromEndOfListCase{
		{
			"case 1",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
	}

	for _, c := range cases {
		createRemoveNthNodeFromEndOfListTestCase(t, c)
	}

}
