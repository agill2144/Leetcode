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
        newHead := dfs(curr.Next)
        if curr.Next != nil {
            curr.Next.Next = curr
            curr.Next = nil
        } else {
            return curr
        }
        return newHead
    }
    return dfs(head)
}

// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {return head}
//     curr := head
//     var prev *ListNode
//     for curr != nil {
//         next := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = next
//     }
//     return prev
// }