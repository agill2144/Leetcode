/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    dummy := &ListNode{}
    tail := dummy
    for l1 != nil || l2 != nil {
        l1Val := 0
        l2Val := 0
        if l1 != nil {l1Val = l1.Val; l1=l1.Next}
        if l2 != nil {l2Val = l2.Val; l2=l2.Next}
        sum := carry + l1Val + l2Val
        tail.Next = &ListNode{Val:sum%10}
        tail = tail.Next
        carry = sum / 10
    }
    if carry != 0 {
        tail.Next = &ListNode{Val: carry}
        carry = 0
    }
    return dummy.Next
}