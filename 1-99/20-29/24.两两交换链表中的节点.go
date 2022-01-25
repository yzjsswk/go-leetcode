package goleetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

/*
 * 模拟
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{0, head}
	i := h
	for i.Next != nil && i.Next.Next != nil {
		j, k := i.Next, i.Next.Next
		tmp := k.Next
		k.Next = j
		i.Next = k
		j.Next = tmp
		i = j
	}
	return h.Next
}

// @lc code=end
