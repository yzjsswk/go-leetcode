package goleetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

/*
 * 如果一个一个链表依次合并，会导致答案链上的指针移动过多次数(每次都要从头开始再移动一遍)
 * 考虑用一个优先队列，每次拿所有链表中头上最小的那个节点进行合并，这样答案链上的指针只需要移动一趟
 *
 * 具体实现时，优先队列中直接存放指针，拿指针的值比较，每次只需要把堆顶指针更新为Next，如果为空了才出堆
 * (或者先pop，next不为空再push回去)
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	queue := Heap{}
	for i := 0; i < l; i++ {
		if lists[i] != nil {
			queue.Push(lists[i])
		}
	}
	heap.Init(&queue)
	ans := &ListNode{0, nil}
	tail := ans
	for queue.Len() != 0 {
		tail.Next = queue[0]
		tail = tail.Next
		queue[0] = queue[0].Next
		if queue[0] == nil {
			heap.Remove(&queue, 0)
		} else {
			heap.Fix(&queue, 0)
		}
	}
	return ans.Next
}

type Heap []*ListNode

func (h Heap) Len() int            { return len(h) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h *Heap) Push(x interface{}) { (*h) = append((*h), x.(*ListNode)) }
func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return x
}

// @lc code=end
