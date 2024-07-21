/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    seen := map[*ListNode]struct{}{}
    curr := head
    for curr != nil {
        if _, ok := seen[curr]; ok {return true}
        seen[curr] = struct{}{}
        curr = curr.Next
    }
    return false
}