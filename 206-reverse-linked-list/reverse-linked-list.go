/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var dfs func(curr, prev *ListNode)
    var newHead *ListNode
    dfs = func(curr, prev *ListNode) {
        // base
        if curr == nil {
            newHead = prev
            return
        }

        // logic
        dfs(curr.Next, curr)
        curr.Next = prev
        if prev != nil {
            prev.Next = nil
        }
    }
    dfs(head, nil)
    return newHead
}