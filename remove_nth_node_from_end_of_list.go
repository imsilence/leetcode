package leetcode

/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	index := length - n
	switch {
	case index == 0:
		return head.Next
	default:
		current = head
		index--
		for index > 0 {
			index--
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	return head
}
