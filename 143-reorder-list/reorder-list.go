/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {return}
    // find middle
    var prev *ListNode
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    // split into 2
    prev.Next = nil
    // reverse 2nd list
    slow = reverse(slow)
    // take 2 ptrs
    p1 := head
    p2 := slow
    // collect node from left and then from right
    // use sentinel node with a tail to collect nodes
    // until both LL ptrs are done   
    dummy := &ListNode{Val: 0}
    tail := dummy
    for p1 != nil || p2 != nil {
        if p1 != nil {
            tail.Next = p1
            p1 = p1.Next
            tail = tail.Next
        }
        if p2 != nil {
            tail.Next = p2
            p2 = p2.Next
            tail = tail.Next
        }
    }
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
