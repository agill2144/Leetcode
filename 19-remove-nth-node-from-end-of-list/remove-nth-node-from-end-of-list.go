/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return head}
    slow := head
    fast := head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    var prev *ListNode
    for fast != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next
    }

    if slow == head {
        return head.Next
    }

    if prev != nil {
        prev.Next = slow.Next
    }
    slow.Next = nil
    return head
}