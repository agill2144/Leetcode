/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    if head == nil {return head}
    head = reverse(head)
    carry := 1
    curr := head
    tail := head
    for curr != nil {
        sum := curr.Val + carry
        curr.Val = sum % 10
        carry = sum / 10
        tail = curr
        curr = curr.Next
    }
    if carry != 0 {
        tail.Next = &ListNode{Val: carry}
        carry = 0
    }
    return reverse(head)
}

func reverse(head *ListNode) *ListNode {
    if head == nil  {return head}
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}