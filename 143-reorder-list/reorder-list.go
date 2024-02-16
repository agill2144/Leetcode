/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil {return}
    slow := head
    fast := head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    p1 := head
    p2 := reverse(slow.Next)
    slow.Next = nil
    
    /*
        p1
        1 - 2 - 3
        |
        4 - 5 - 6
        p2
    */
    for p1 != nil && p2 != nil {
        p1Next := p1.Next
        p2Next := p2.Next
        p1.Next = p2
        p2.Next = p1Next
        p1 = p1Next
        p2 = p2Next
    }
    
    
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}