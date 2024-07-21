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
    second := slow
    second = reverse(second)
    

    p1, p2 := head, second
    for p1 != nil && p2 != nil {
        if p1.Val != p2.Val {return false}
        p1 = p1.Next
        p2 = p2.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
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