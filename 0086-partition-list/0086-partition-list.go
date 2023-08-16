/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {return head}
    
    left := &ListNode{Val: -1}
    leftTail := left
    right := &ListNode{Val: -1}
    rightTail := right

    curr := head
    for curr != nil {
        next := curr.Next
        if curr.Val < x {
            curr.Next = nil
            leftTail.Next = curr
            leftTail = leftTail.Next
        } else {
            curr.Next = nil
            rightTail.Next = curr
            rightTail = rightTail.Next
        }
        curr = next
    }
    // speed up from dummy head nodes
    left = left.Next
    right = right.Next
    // if there is a left side to connect
    if left != nil {
        leftTail.Next = right
        return left
    }
    return right
}