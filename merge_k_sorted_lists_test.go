package leetcode

import "testing"

type MergeKSortedListsCase struct {
	name   string
	lists  []*ListNode
	result *ListNode
}

func (c *MergeKSortedListsCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createMergeKSortedListsTestCase(t *testing.T, c *MergeKSortedListsCase) {
	t.Helper()
	result := mergeKLists(c.lists)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestMergeKLists(t *testing.T) {
	cases := []*MergeKSortedListsCase{
		{
			"case 1",
			[]*ListNode{
				{1, &ListNode{2, &ListNode{4, nil}}},
				{1, &ListNode{3, &ListNode{4, nil}}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			"case 2",
			[]*ListNode{
				{1, &ListNode{4, &ListNode{5, nil}}},
				{1, &ListNode{3, &ListNode{4, nil}}},
				{2, &ListNode{6, nil}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
		},
		{
			"case 3",
			[]*ListNode{},
			nil,
		},
	}

	for _, c := range cases {
		createMergeKSortedListsTestCase(t, c)
	}

}
