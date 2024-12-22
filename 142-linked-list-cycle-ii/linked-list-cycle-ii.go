/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {break}
    }
    if fast == nil || fast.Next == nil {return nil}

    fast = head
    for slow != fast {
        fast = fast.Next
        slow = slow.Next
    }
    return fast
}