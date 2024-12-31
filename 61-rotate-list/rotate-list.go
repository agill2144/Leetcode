/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {return head}
    n := 0
    curr := head
    ogTail := head
    for curr != nil {
        n++
        ogTail = curr
        curr = curr.Next
    }
    k %= n
    if k == 0 {return head}
    curr = head
    tail := head
    for i := 0; i < n-k; i++ {
        tail = curr
        curr = curr.Next        
    }
    newHead := tail.Next
    tail.Next = nil
    ogTail.Next = head
    head = newHead
    return head    
}