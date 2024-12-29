/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    // we are creating k groups
    // of size m each
    m := n/k
    r := n%k
    curr = head
    out := make([]*ListNode, k)
    i := 0
    for i < len(out) {
        currSize := 0
        h := curr
        tail := curr
        for currSize != m && curr != nil {
            tail = curr
            curr = curr.Next
            currSize++
        }
        if r != 0 {
            tail = curr; curr=curr.Next
            r--
        }
        if tail != nil {tail.Next = nil}
        out[i] = h
        i++
    }
    return out
}