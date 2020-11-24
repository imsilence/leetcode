package leetcode

import "testing"

type AddTwoNumbersCase struct {
	name   string
	left   *ListNode
	right  *ListNode
	result *ListNode
}

func (c *AddTwoNumbersCase) ok(r *ListNode) bool {
	return c.result.Equals(r)
}

func createAddTwoNumberTestCase(t *testing.T, c *AddTwoNumbersCase) {
	t.Helper()
	result := addTwoNumbers(c.left, c.right)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []*AddTwoNumbersCase{
		{
			"case 1",
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}

	for _, c := range cases {
		createAddTwoNumberTestCase(t, c)
	}

}
