package leetcode

/*
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	var head, current *ListNode
	indexs := append([]*ListNode{}, lists...)

	next := func() (int, bool) {
		var (
			minIndex int = 0
			minNode  *ListNode
		)
		for i, node := range indexs {
			if node == nil {
				continue
			}
			if minNode == nil || minNode.Val > node.Val {
				minNode = node
				minIndex = i
			}
		}

		if minNode == nil {
			return 0, false
		}
		indexs[minIndex] = indexs[minIndex].Next
		return minNode.Val, true
	}

	appendList := func(v int) {
		if head == nil {
			head = &ListNode{v, nil}
			current = head
		} else {
			current.Next = &ListNode{v, nil}
			current = current.Next
		}
	}

	for {
		v, ok := next()
		if !ok {
			break
		}
		appendList(v)
	}

	return head
}
