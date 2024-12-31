/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {return head}
    if head == nil || head.Next == nil {return head}
    var pt *ListNode
    curr := head
    size := 1
    var start *ListNode
    for curr != nil {
        next := curr.Next
        if size == 1 {start = curr}
        if size == k {
            // disconnect this k size LL 
            // from next and prev nodes
            curr.Next = nil
            if pt != nil {pt.Next = nil}
            // now reverse this isolated / disconnected LL
            nh, nt := reverse(start)
            // connect prevTail to newHead
            if pt != nil {
                pt.Next = nh
            } else {
                head = nh
            }
            // connect new tail to next node that is outside of this LL
            nt.Next = next
            // now we have connected the reversed version of k size LL
            // prevTail now moves to our newTail for the next k size grp
            pt = nt
            curr = next
            start = nil
            size = 1
            continue
        }
        size++
        curr = next
    }
    return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
    if head == nil || head.Next == nil {return head, head}
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