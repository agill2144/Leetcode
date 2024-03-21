/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
