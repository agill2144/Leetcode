// optimized : o(n)
// space: o(1)
func swapNodes(head *ListNode, k int) *ListNode {
    	if head == nil || head.Next == nil || k < 1 {
		return head
	}
	current := head
	var front *ListNode
	var kthFromEnd *ListNode
	size := 0
	for current != nil {
		size++
		if kthFromEnd != nil {
			kthFromEnd = kthFromEnd.Next
		}
		if size == k {
			front = current
			kthFromEnd = head
		}
		current = current.Next
	}
	front.Val , kthFromEnd.Val = kthFromEnd.Val, front.Val
	return head
}