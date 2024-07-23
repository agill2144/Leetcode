/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {return head}
    var prev *ListNode
    curr := head
    for curr != nil {
        count := 1
        t := curr
        for count != k && t != nil {
            t = t.Next
            count++
        }
        if t == nil {break}
        next := t.Next
        if prev != nil {prev.Next = nil}
        t.Next = nil
        newHead, newTail := reverse(curr)
        if prev != nil {prev.Next = newHead} else {head = newHead}
        newTail.Next = next
        prev = newTail
        curr = next
    }
    return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
    if head == nil || head.Next == nil {return head, head}
    newTail := head
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev, newTail
}