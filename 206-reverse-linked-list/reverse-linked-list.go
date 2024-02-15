/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
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


 // dfs
// func reverseList(head *ListNode) *ListNode {
//     var dfs func(curr, prev *ListNode)
//     var newHead *ListNode
//     dfs = func(curr, prev *ListNode) {
//         // base
//         if curr == nil {
//             newHead = prev
//             return
//         }

//         // logic
//         dfs(curr.Next, curr)
//         curr.Next = prev
//         if prev != nil {
//             prev.Next = nil
//         }
//     }
//     dfs(head, nil)
//     return newHead
// }

func reverseList(head *ListNode) *ListNode {
    // base
    if head == nil {return head}

    // logic
    result := reverseList(head.Next)
    if head.Next != nil {
        head.Next.Next = head
        head.Next = nil
    } else {
        return head
    }
    return result

}
