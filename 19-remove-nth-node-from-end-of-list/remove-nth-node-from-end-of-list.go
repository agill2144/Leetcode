/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return head}
    p1, p2 := head, head
    for i := 0; i < n; i++ {
        p2 = p2.Next
    }
    var prev *ListNode
    for p2 != nil {
        prev = p1
        p1 = p1.Next
        p2 = p2.Next
    }
    if prev == nil {
        newHead := head.Next
        head.Next = nil
        head = newHead
    } else {
        prev.Next = p1.Next
        p1.Next = nil
    }
    return head
}