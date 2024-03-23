/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {return}
    
    // 1. split list in 2 halves
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 2. reverse the 2nd half
    head2 := slow.Next
    slow.Next = nil
    head2 = reverse(head2)

    // 3. using 2 ptrs, connect the 2 lists
    c1 := head
    c2 := head2
    for c1 != nil && c2 != nil {
        c1Next := c1.Next
        c2Next := c2.Next
        c1.Next = c2
        c2.Next = c1Next
        c1 = c1Next
        c2 = c2Next
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