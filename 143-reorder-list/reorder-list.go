/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {return}
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    p1 := head
    p2 := slow.Next
    slow.Next = nil
    p2 = reverse(p2)
    // 1-2
    // 5-4-3
    var tail *ListNode
    for p1 != nil && p2 != nil {
        p1Next := p1.Next
        p2Next := p2.Next
        p1.Next = p2
        p2.Next = p1Next
        p1 = p1Next
        p2 = p2Next
    }
    for p2 != nil {
        tail.Next = p2
        tail = tail.Next
        p2 = p2.Next
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
