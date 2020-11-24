package leetcode

import "testing"

type SwapNodesInPairsCase struct {
	name   string
	list   *ListNode
	result *ListNode
}

func (c *SwapNodesInPairsCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createSwapNodesInPairsTestCase(t *testing.T, c *SwapNodesInPairsCase) {
	t.Helper()
	result := swapPairs(c.list)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestSwapPairs(t *testing.T) {
	cases := []*SwapNodesInPairsCase{
		{
			"case 1",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{2, &ListNode{1, &ListNode{4, nil}}},
		},
		{
			"case 2",
			nil,
			nil,
		},
		{
			"case 3",
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			"case 4",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		},
	}

	for _, c := range cases {
		createSwapNodesInPairsTestCase(t, c)
	}

}
