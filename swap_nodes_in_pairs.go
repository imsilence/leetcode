package leetcode

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]


提示：

链表中节点的数目在范围 [0, 100] 内
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var result, prev *ListNode

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head

		if result == nil {
			result = next
		} else {
			prev.Next = next
		}

		prev = head
		head = head.Next
	}

	return result
}
