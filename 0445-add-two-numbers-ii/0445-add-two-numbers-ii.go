/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1 = reverse(l1)
    l2 = reverse(l2)
    
    var out *ListNode
    carry := 0
    
    l1Ptr, l2Ptr := l1, l2
    for l1Ptr != nil || l2Ptr != nil {
        x := 0
        if l1Ptr != nil {
            x = l1Ptr.Val
            l1Ptr = l1Ptr.Next
        }
        y := 0
        if l2Ptr != nil {
            y = l2Ptr.Val
            l2Ptr = l2Ptr.Next
        }
        sum := x+y+carry
        nodeVal := sum%10
        carry = sum/10
        if sum >= 10 {
            nodeVal = sum % 10
            carry = sum / 10
        } else {
            carry = 0
        }
        newHead := &ListNode{Val: nodeVal}
        newHead.Next = out
        out = newHead
    }
    if carry != 0 {
        newHead := &ListNode{Val: carry}
        newHead.Next = out
        out = newHead        
    }
    return out
    
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev= curr
        curr = next
    }
    return prev
}