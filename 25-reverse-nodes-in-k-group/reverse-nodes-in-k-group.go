/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k <= 1 {return head}
    size := 1
    curr := head
    var start *ListNode
    var prevTail *ListNode
    for curr != nil {
        next := curr.Next
        if size == 1 {start = curr}
        if size == k {
            if prevTail != nil {prevTail.Next = nil}
            curr.Next = nil
            nh, nt := reverse(start)
            if prevTail != nil {
                prevTail.Next = nh
            } else {
                head = nh
            }
            nt.Next = next
            prevTail = nt
            size = 1
            curr = next
            continue
        }
        size++
        curr = next
    }
    return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
    if head == nil || head.Next == nil {
        return head, head
    }
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev, head
}