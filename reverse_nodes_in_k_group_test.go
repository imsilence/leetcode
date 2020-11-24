package leetcode

import "testing"

type RevereNodesInKGroupCase struct {
	name   string
	list   *ListNode
	k      int
	result *ListNode
}

func (c *RevereNodesInKGroupCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createRevereNodesInKGroupTestCase(t *testing.T, c *RevereNodesInKGroupCase) {
	t.Helper()
	result := reverseKGroup(c.list, c.k)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestReverseKGroup(t *testing.T) {
	cases := []*RevereNodesInKGroupCase{
		{
			"case 1",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			"case 2",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			3,
			&ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}

	for _, c := range cases {
		createRevereNodesInKGroupTestCase(t, c)
	}

}
