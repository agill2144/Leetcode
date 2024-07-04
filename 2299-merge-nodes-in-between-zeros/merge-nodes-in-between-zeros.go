/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    if head == nil {return head}
    if head.Next == nil {
        if head.Val == 0 {return nil}
        return head
    }
    curr := head
    var prev *ListNode // prev summed node
    for curr != nil && curr.Next != nil {
        if curr.Val == 0 {
            start := curr
            sum := 0
            curr = curr.Next
            for curr != nil && curr.Val != 0 {
                sum += curr.Val
                curr = curr.Next
            }
            newNode := &ListNode{Val: sum}
            if start == head {
                head = newNode
            }
            if prev != nil {
                prev.Next = newNode
            }
            prev = newNode
        }
    }
    return head
}