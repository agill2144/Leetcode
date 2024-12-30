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
            // disconnect this group from next grp
            curr.Next = nil
            // reverse this grp
            nh, nt := reverse(start)
            // if there was a prev grp ( pt = prevTail )
            // then connect pt with newHead (nh)
            if pt != nil {
                pt.Next = nh
            } else {
                // if there was no pt, it means this is 1st k size grp
                // nh is the head node now
                head = nh
            }
            // connect this grp newTail with next node
            nt.Next = next
            // move pt ptr to our newTail (nt)
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