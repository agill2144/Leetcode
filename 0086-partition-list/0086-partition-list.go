/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {return head}
    left := &ListNode{Val: 0}
    leftTail := left
    
    right := &ListNode{Val:0}
    rightTail := right
    
    curr := head
    for curr != nil {
        next := curr.Next
        if curr.Val < x {
            leftTail.Next = curr
            leftTail = leftTail.Next
        } else {
            rightTail.Next = curr
            rightTail = rightTail.Next                
        }
        curr.Next = nil
        curr = next
    }
    left = left.Next
    right = right.Next
    
    if left == nil {
        return right
    }
    leftTail.Next = right
    return left
}