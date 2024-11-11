/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {return true}
    slow := head
    fast := head
    var prev *ListNode
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    h1 := head
    h2 := slow
    h2 = reverse(h2)
    p1, p2 := h1, h2
    for p1 != nil && p2 != nil {
        if p1.Val != p2.Val {return false}
        p1 = p1.Next
        p2 = p2.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
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