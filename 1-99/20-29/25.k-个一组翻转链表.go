package goleetcode

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

/*
 * 《模拟》
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseKGroup(head *ListNode, k int) *ListNode {
	h := &ListNode{0, head}
	ok_tail := h
	flag := true
	for flag && ok_tail.Next != nil {
		i, j := ok_tail, ok_tail.Next
		cnt := 0
		for cnt < k && j != nil {
			tmp := j.Next
			j.Next = i
			i = j
			j = tmp
			cnt++
		}
		if cnt == k {
			tmp := ok_tail.Next
			tmp.Next = j
			ok_tail.Next = i
			ok_tail = tmp
		} else {
			for i != ok_tail {
				tmp := i.Next
				i.Next = j
				j = i
				i = tmp
				flag = false
			}
		}
	}
	return h.Next
}

// @lc code=end
