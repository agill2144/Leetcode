/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return nil}
    slow := head
    fast := head
    var init *ListNode 
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            init = slow
            break
        }
    }
    if init == nil {return nil}
        
    current := head
    for current != init {
        current = current.Next
        init = init.Next
    }
    return current
}