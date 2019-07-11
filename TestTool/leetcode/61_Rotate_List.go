package main

func main() {
	
}

// 自写的100% + 100%
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	current, before := head, head
	length := 0
	for current != nil {
		current = current.Next
		length++
	}

	k = k%length

	for current = head; k>0; k-- {
		current = current.Next
	}

	for current.Next != nil {
		before, current = before.Next, current.Next
	}

	current.Next = head
	newHead := before.Next
	before.Next = nil

	return newHead
}