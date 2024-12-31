/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    out := make([]*ListNode, k)
    if head == nil {return out}
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    grpSize := n / k
    rem := n % k
    curr = head
    outPtr := 0
    for curr != nil {
        start := curr
        tail := curr
        size := 0
        for size != grpSize && curr != nil {
            size++
            tail = curr
            curr = curr.Next
        }
        if rem != 0 && curr != nil {
            tail = curr
            curr = curr.Next
            rem--
        }
        tail.Next = nil
        out[outPtr] = start
        outPtr++
    }
    return out
}