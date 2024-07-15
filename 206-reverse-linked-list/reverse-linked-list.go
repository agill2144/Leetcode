/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// recursive
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var dfs func(curr *ListNode, prev *ListNode) *ListNode
    dfs = func(curr *ListNode, prev *ListNode) *ListNode {
        // base
        if curr == nil {return prev}
        // logic
        next := curr.Next
        curr.Next = prev
        return dfs(next, curr)
    }
    return dfs(head, nil)
}

/*
    approach: classic
    - point to prev
    - prev at the end becomes the new 
    time = o(n)
    space = o(1)
*/
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