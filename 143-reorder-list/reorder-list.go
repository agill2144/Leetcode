/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {return}
    var prev *ListNode    
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    head2 := slow
    prev = nil
    curr := head2
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    head2 = prev

    dummy := &ListNode{Val:0}
    tail := dummy
    p1 := head
    p2 := head2
    for p1 != nil || p2 != nil {
        if p1 != nil {
            tail.Next = p1
            tail = tail.Next
            p1 = p1.Next
        }
        if p2 != nil {
            tail.Next = p2
            tail = tail.Next
            p2 = p2.Next
        }
    }

}