/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    if k <= 1 {
        return []*ListNode{head}
    }
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    sizeOfEach := n/k
    curr = head
    out := []*ListNode{}
    start := curr
    size := 0
    for curr != nil {
        size++
        next := curr.Next
        if size < sizeOfEach || (len(out) < n%k && size < sizeOfEach+1) {
            curr = next
            continue
        }
        
        curr.Next = nil
        out = append(out, start)
        start = next
        curr = next
        size = 0
    }
    for len(out) != k {
        out = append(out, nil)
    }
    return out
}