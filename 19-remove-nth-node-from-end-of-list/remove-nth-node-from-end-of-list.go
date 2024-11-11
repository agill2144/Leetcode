/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return head}
    p1 := head
    p2 := head
    for i := 0; i < n; i++ {
        p2 = p2.Next
    }
    var prev *ListNode
    for p2 != nil {
        prev = p1
        p1 = p1.Next
        p2 = p2.Next
    }
    if prev != nil {
        prev.Next = p1.Next
        p1.Next = nil
    } else if prev == nil {
        // node to delete is head node
        newHead := head.Next
        head.Next = nil
        head = newHead
    }
    return head
}