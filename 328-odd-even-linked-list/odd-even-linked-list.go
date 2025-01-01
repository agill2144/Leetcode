/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    odd := &ListNode{Val:0}
    even := &ListNode{Val: 0}
    ot := odd
    et := even
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = nil
        // first node is considered odd, 
        // and the second node is even, and so on
        ot.Next = curr
        ot = ot.Next
        curr = next
        if curr != nil {
            next = curr.Next
            curr.Next = nil
            et.Next = curr
            et = et.Next
            curr = next
        }
    }
    odd = odd.Next
    even = even.Next
    ot.Next = even
    return odd
}