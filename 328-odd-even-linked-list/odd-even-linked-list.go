/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    d1 := &ListNode{Val:-1}; t1 := d1
    d2 := &ListNode{Val:-1}; t2 := d2
    
    curr := head
    for curr != nil {
        // The first node is considered odd, and the second node is even, and so on.
        next := curr.Next
        curr.Next = nil
        t1.Next = curr
        t1 = t1.Next
        curr = next
        if curr != nil {
            next = curr.Next
            curr.Next = nil
            t2.Next = curr
            t2 = t2.Next
            curr = next
        }
    }
    d1 = d1.Next
    d2 = d2.Next
    t1.Next = d2
    return d1
    
}