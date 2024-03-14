/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    dummy.Next = head
    curr := head
    rSum := 0
    sumToNode := map[int]*ListNode{0:dummy}
    
    for curr != nil {
        next := curr.Next
        rSum += curr.Val
        prev, ok := sumToNode[rSum]
        if ok {
            start := prev.Next
            for start != curr {
                // need to clean up the map :/ 
                for s, n := range sumToNode {
                    if n == start {
                        delete(sumToNode, s)
                        break
                    }
                }
                start = start.Next
            }
            prev.Next = next
            curr.Next = nil
        } else {
            sumToNode[rSum] = curr
        }
        curr = next
    }
    return dummy.Next
}