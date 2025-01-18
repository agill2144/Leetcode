/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    odd := &ListNode{}
    ot := odd
    even := &ListNode{}
    et := even
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = nil
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
    even = even.Next
    odd = odd.Next
    ot.Next = even
    return odd   
}