/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {return nil}
    o := &ListNode{Val: -1}; ot := o
    e := &ListNode{Val: -1}; et := e
    i := 0
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = nil
        if i % 2 == 0 {
            et.Next = curr
            et = et.Next
        } else {
            ot.Next = curr
            ot = ot.Next
        }
        curr = next
        i++
    }
    o = o.Next
    e = e.Next
    et.Next = o
    return e
}