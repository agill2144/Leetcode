/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    rev := reverse(head)
    carry := 1
    curr := rev
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        sum := curr.Val + carry
        curr.Val = sum % 10
        carry = sum / 10
        prev = curr
        curr = next
    }
    if carry != 0 {
        prev.Next = &ListNode{Val: carry}
    }
    return reverse(rev)
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