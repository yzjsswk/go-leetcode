package goleetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

/*
 * 一次遍历：弄两个指针前后隔n个节点，后一个指针到最后的时候，
 * 前一个指针的位置就是要进行删除操作的位置
 *
 * [1] 前一个指针i实际上最后指向的是要删的节点的前一个节点，
 * 如果要删的节点是第一个节点，他前面就没有节点了
 * 所以先在第一个节点前加一个节点，使下面的逻辑统一
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{0, head} // [1]
	i, j := h, h
	for ; n >= 0; n-- {
		j = j.Next
	}
	for j != nil {
		i, j = i.Next, j.Next
	}
	i.Next = i.Next.Next
	return h.Next
}

// @lc code=end
