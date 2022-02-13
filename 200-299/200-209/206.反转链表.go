package goleetcode

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := &ListNode{0, head}
	i, j := h, h.Next
	for j != nil {
		tmp := j.Next
		j.Next = i
		i = j
		j = tmp
	}
	h.Next.Next = nil
	return i
}

// @lc code=end
