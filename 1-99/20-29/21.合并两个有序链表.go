package goleetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

/*
 * 模拟，注意考虑边界情况，添加头节点可以统一逻辑
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{-101, list1}
	i, j := h, list2
	for j != nil {
		for i.Next != nil && i.Next.Val < j.Val {
			i = i.Next
		}
		if i.Next == nil {
			i.Next = j
			break
		}
		tmp := j.Next
		j.Next = i.Next
		i.Next = j
		j = tmp
	}
	return h.Next
}

// @lc code=end
