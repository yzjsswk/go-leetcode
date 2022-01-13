package goleetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	tail := sum
	i, j := l1, l2
	c := 0
	for {
		a, b := 0, 0
		if i == nil && j == nil {
			break
		}
		if i != nil {
			a, i = i.Val, i.Next
		}
		if j != nil {
			b, j = j.Val, j.Next
		}
		tail.Next = &ListNode{(a + b + c) % 10, nil}
		tail = tail.Next
		c = (a + b + c) / 10
	}
	if c == 1 {
		tail.Next = &ListNode{1, nil}
	}
	return sum.Next
}

// @lc code=end
