/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := map[*ListNode]struct{}{}
    curr := head
    for curr != nil {
        if _, ok := m[curr]; ok {return true}
        m[curr] = struct{}{}
        curr = curr.Next
    }
    return false
}