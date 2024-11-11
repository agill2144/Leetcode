/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    set := map[*ListNode]bool{}
    curr := head
    for curr != nil {
        if set[curr] {return curr}
        set[curr] = true
        curr = curr.Next
    }
    return nil
}