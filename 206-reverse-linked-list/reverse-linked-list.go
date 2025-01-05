/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var dfs func(curr *ListNode) *ListNode
    dfs = func(curr *ListNode) *ListNode {
        // base
        if curr == nil {return nil}

        // logic
        tail := dfs(curr.Next)
        if curr.Next == nil {
            return curr
        } else {
            curr.Next.Next = curr
            curr.Next = nil
            
        }
        return tail
    }
    return dfs(head)
}