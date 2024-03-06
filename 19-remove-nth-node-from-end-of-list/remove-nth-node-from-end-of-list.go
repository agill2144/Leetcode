/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prev *ListNode
    p1, p2 := head, head
    for i := 0; i < n && p2 != nil ; i++ {
        p2 = p2.Next
    }
    if p2 == nil {
        // remove head
        return p1.Next
    }

    for p2 != nil {
        prev = p1
        p1 = p1.Next
        p2 = p2.Next
    }
    prev.Next = p1.Next
    p1.Next = nil
    return head
}