/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    curr := head
    var prev *ListNode
    for curr != nil && curr.Next != nil {
        adj := curr.Next
        next := curr.Next.Next
        
        curr.Next = next
        adj.Next = curr
        if curr == head {
            head = adj
        }
        
        if prev != nil {
            prev.Next = adj
        }
        prev = curr
        curr = next
    }
    return head
}