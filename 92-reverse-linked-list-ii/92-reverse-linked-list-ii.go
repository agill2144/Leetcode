/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseInPlace(head *ListNode) (*ListNode, *ListNode) {
// 	if head == nil {
// 		return nil,nil
// 	}
// 	current := head
// 	var tail *ListNode
// 	var prev *ListNode
// 	for current != nil {
// 		next := current.Next
// 		current.Next = prev
// 		if current.Next == nil {
// 			tail = current
// 		}
// 		prev = current
// 		current = next

// 	}
// 	// head, tail of reversed
// 	return prev, tail
// }


// func reverseBetween(head *ListNode, left, right int) *ListNode {
// 	if head == nil || left < 0 || right < 0 || head.Next == nil {
// 		return head
// 	}
// 	var start *ListNode
// 	var end *ListNode
// 	current := head
// 	currentPos := 1
// 	var toBeReversed = &ListNode{Val: 0}
// 	currentToBeReversed := toBeReversed
// 	for current != nil {
// 		next := current.Next
// 		if currentPos == left-1 {
// 			start = current
// 		}
// 		if currentPos == right+1 {
// 			end = current
// 		}
// 		if currentPos >= left && currentPos <= right {
// 			currentToBeReversed.Next = &ListNode{Val: current.Val}
// 			currentToBeReversed = currentToBeReversed.Next
// 		}
// 		current = next
// 		currentPos++
// 	}

// 	toBeReversed = toBeReversed.Next
// 	revHead, revTail := reverseInPlace(toBeReversed)
// 	if start == nil {
// 		// i.e left was a 1
// 		head = revHead
// 	} else {
// 		start.Next = revHead
// 	}
// 	revTail.Next = end
// 	return head
// }


func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil || left < 1 || right < 1 || head.Next == nil {
		return head
	}
	var (
		leftAttachment *ListNode
		rightAttachment *ListNode
		subListTail *ListNode
		subListHead *ListNode
		prev *ListNode
	)
	current := head
	currentPos := 1
	for current != nil {
		next := current.Next
		if left-1 == currentPos {
			leftAttachment = current
		}
		if right+1 == currentPos {
			rightAttachment = current
		}
		if currentPos >= left && currentPos <= right {
			if currentPos == left {
				subListTail = current
			}
			if currentPos == right {
				subListHead = current
			}
			current.Next = prev
		}
		prev = current
		current = next
		currentPos++
	}
	if leftAttachment == nil {
		head = subListHead
	} else {
		leftAttachment.Next = subListHead
	}
	subListTail.Next = rightAttachment
	return head
}