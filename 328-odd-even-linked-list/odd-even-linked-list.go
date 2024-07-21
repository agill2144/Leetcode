/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    dummy := &ListNode{Val: 0}
    eTail := dummy
    for curr != nil && curr.Next != nil {
        next := curr.Next
        nextNext := next.Next

        eTail.Next = next
        curr.Next = nil
        eTail = eTail.Next
        eTail.Next = nil

        curr.Next = nextNext
        if nextNext != nil {
            curr = nextNext
        }
    }
    curr.Next = dummy.Next
    return head
}