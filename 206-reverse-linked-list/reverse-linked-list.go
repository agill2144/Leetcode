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
        if curr.Next == nil {
            // we are at the tail end of original LL
            // the tail is the new head, therefore return it
            return curr
        } else {
            // next node's next ptr should point to me!
            curr.Next.Next = curr
            curr.Next = nil
        }
        return tail
    }
    return dfs(head)
}