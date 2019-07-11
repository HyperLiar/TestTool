package main

func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode

	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}

	return prev
}

// 正常思维
func reverseList1(head *ListNode) *ListNode {
	current := head
	var prev *ListNode

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}