/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    start := head
    var startPrev *ListNode
    count := 1
    
    for curr != nil {
        next := curr.Next
        if count == k {
            
            // reverse from start to curr(including curr)
            curr.Next = nil
            newHead, newTail := reverse(start)
            // connect new tail with rest of the LL
            newTail.Next = next

            // connect prev LL with the new head 
            if startPrev != nil {startPrev.Next = newHead}

            if newTail == head { head = newHead }
            
            startPrev = newTail
            curr = next
            start = next
            count = 1
            
            
        } else {
            count++
            curr = next
        }
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