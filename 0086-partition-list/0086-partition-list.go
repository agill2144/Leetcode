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
        if curr.Val >= x {
            rightTail.Next = &ListNode{Val: curr.Val}
            rightTail = rightTail.Next
        } else {
            leftTail.Next = &ListNode{Val: curr.Val}
            leftTail = leftTail.Next
        }
        curr = curr.Next
    }
    left = left.Next
    right = right.Next
    if left != nil {
        leftTail.Next = right
        return left
    }
    return right
}