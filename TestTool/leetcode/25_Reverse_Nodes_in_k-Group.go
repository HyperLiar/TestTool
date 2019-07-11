package main

import "fmt"

/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list.
If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, 0)

	var newHead *ListNode
	before := &ListNode{}

	for current := head; ; current = current.Next {
		if len(stack) == k {
			if newHead == nil {
				newHead = stack[len(stack)-1]
			}
			for i := k - 1; i >= 0; i-- {
				before.Next = stack[i]
				before = before.Next
			}
			before.Next = current
			stack = []*ListNode{}
		}
		stack = append(stack, current)

		if current == nil {
			break
		}
	}
	if newHead == nil {
		return head
	}

	return newHead
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	beg, end := reverseList(head.Next)
	head.Next = nil
	end.Next = head
	return beg, head
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	next := head
	for i := 0 ; i < k - 1 ; i++ {
		if next == nil {
			return head
		}
		next = next.Next
	}
	if next == nil {
		return head
	}
	// fmt.Println(head.Val)
	second := next.Next
	next.Next = nil
	second = reverseKGroup(second, k)
	first, end := reverseList(head)
	end.Next = second
	return first
}

/**
[1,2,3,4,5]
2
 */
func main() {
	head := &ListNode{}
	current := head

	for i := 1; i < 3; i++ {
		current.Val = i
		current.Next = &ListNode{}
		current = current.Next
	}

	current = head
	for current.Next != nil {
		fmt.Println(current.Val)
		current = current.Next
	}

	fmt.Println("reverse:")
	head = reverseKGroup(head, 2)

	idx := 0
	for head.Next != nil && idx < 10 {
		fmt.Println(head.Val)
		head = head.Next
		idx++
	}
}
