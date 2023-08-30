/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    left := &ListNode{Val: 0}
    leftTail := left
    right := &ListNode{Val: 0}
    rightTail := right
    
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = nil
        if curr.Val >= x {
            rightTail.Next = curr
            rightTail = rightTail.Next
        } else {
            leftTail.Next = curr
            leftTail = leftTail.Next
        }
        curr = next
    }
    left = left.Next
    right = right.Next
    if left != nil {
        leftTail.Next = right
        return left
    }
    return right
}