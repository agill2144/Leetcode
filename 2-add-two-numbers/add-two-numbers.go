/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    tail := dummy

    carry := 0
    p1, p2 := l1, l2
    for p1 != nil || p2 != nil {
        p1Val := 0
        if p1 != nil {p1Val = p1.Val; p1 = p1.Next}
        p2Val := 0
        if p2 != nil {p2Val = p2.Val; p2 = p2.Next}
        sum := p1Val + p2Val + carry
        digit := sum % 10
        carry = sum / 10
        newNode := &ListNode{Val: digit}
        tail.Next = newNode
        tail = tail.Next 
    }
    if carry != 0 {
        tail.Next = &ListNode{Val: carry}
        carry = 0
    }
    return dummy.Next

}
