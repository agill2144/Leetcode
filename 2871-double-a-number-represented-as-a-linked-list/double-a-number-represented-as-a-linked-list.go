/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head == nil {return head}
    revHead := reverse(head)
    carry := 0
    curr := revHead
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        newVal := (curr.Val*2) + carry
        if newVal >= 10 {
            carry = newVal / 10
        } else {
            carry = 0
        }
        curr.Val = newVal%10
        prev = curr
        curr = next
    }
    if carry != 0 {
        newTail := &ListNode{Val: carry}
        prev.Next = newTail
    }
    return reverse(revHead)
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}