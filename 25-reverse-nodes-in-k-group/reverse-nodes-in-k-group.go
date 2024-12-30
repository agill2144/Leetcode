/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 || head == nil || head.Next == nil {return head}
    curr := head
    var pt *ListNode
    var start *ListNode
    size := 1
    for curr != nil {
        next := curr.Next
        if size == 1 {start = curr}
        if size == k {
            curr.Next = nil
            nh, nt := reverse(start)
            if pt != nil {
                pt.Next = nh
            } else {
                head = nh
            }
            nt.Next = next
            pt = nt
            size = 1
            curr = next
            start = nil
            continue
        }
        size++
        curr = next
    }
    return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
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