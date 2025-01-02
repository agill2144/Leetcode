/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(x *ListNode, y *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	tail := dummy
	p1, p2 := x, y
	for p1 != nil || p2 != nil {
		p1Val := math.MaxInt64
		if p1 != nil { p1Val = p1.Val }
		p2Val := math.MaxInt64
		if p2 != nil { p2Val = p2.Val }
		if p1Val < p2Val {
			tail.Next = p1
			tail = tail.Next
			p1 = p1.Next
		} else {
            tail.Next = p2
            tail = tail.Next
            p2 = p2.Next
		}
	}
	return dummy.Next
}