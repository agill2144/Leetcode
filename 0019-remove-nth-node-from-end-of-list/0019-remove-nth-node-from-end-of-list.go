/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return nil}
    p1 := head
    p2 := head
    
    for i := 0; i < n; i++ {
        p1 = p1.Next
    }
    
    var prev *ListNode
    for p1 != nil {
        prev = p2
        p2 = p2.Next
        p1 = p1.Next
    }
    if prev == nil {
        newHead := p2.Next
        p2.Next = nil
        head = newHead
        return head
    }
    prev.Next = p2.Next
    p2.Next = nil
    return head
}