/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var dfs func(curr *ListNode) *ListNode
    dfs = func(curr *ListNode) *ListNode {
        // base
        if curr == nil {return nil}
        // logic
        tail := dfs(curr.Next)
        if curr.Next != nil {curr.Next.Next = curr; curr.Next = nil} else {return curr}
        return tail
    }
    return dfs(head)
}