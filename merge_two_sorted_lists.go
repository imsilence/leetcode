package leetcode

/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, current *ListNode

	appendList := func(v int) {
		if head == nil {
			head = &ListNode{v, nil}
			current = head
		} else {
			current.Next = &ListNode{v, nil}
			current = current.Next
		}
	}

	for l1 != nil && l2 != nil {
		if l1.Val == l2.Val {
			appendList(l1.Val)
			appendList(l1.Val)
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Val > l2.Val {
			appendList(l2.Val)
			l2 = l2.Next
		} else {
			appendList(l1.Val)
			l1 = l1.Next
		}
	}

	for l1 != nil {
		appendList(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		appendList(l2.Val)
		l2 = l2.Next
	}
	return head
}
