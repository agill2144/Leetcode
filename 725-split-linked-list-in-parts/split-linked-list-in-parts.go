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
    // If there are N nodes in the list, and k parts, 
    // then every part has N/k elements, 
    // except the first N%k parts have an extra one.
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
    // its possible k > n
    // then we need to fill output array to match k size with nil LL heads
    for len(out) != k {
        out = append(out, nil)
    }
    return out
}