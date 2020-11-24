package leetcode

import "testing"

type MergeTwoSortedListsCase struct {
	name   string
	left   *ListNode
	right  *ListNode
	result *ListNode
}

func (c *MergeTwoSortedListsCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createMergeTwoSortedListsTestCase(t *testing.T, c *MergeTwoSortedListsCase) {
	t.Helper()
	result := mergeTwoLists(c.left, c.right)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestMergeTwoLists(t *testing.T) {
	cases := []*MergeTwoSortedListsCase{
		{
			"case 1",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}

	for _, c := range cases {
		createMergeTwoSortedListsTestCase(t, c)
	}

}
