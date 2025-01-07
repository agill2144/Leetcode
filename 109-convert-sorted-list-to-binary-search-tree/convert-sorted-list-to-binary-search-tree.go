/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    n := size(head)
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        // logic
        mid := left+(right-left)/2
        leftNode := dfs(left, mid-1)
        if head != nil {
            root := &TreeNode{Val: head.Val}
            head = head.Next
            root.Left = leftNode
            root.Right = dfs(mid+1, right)
            return root
        }
        return nil
    }
    return dfs(0, n-1)
}

func size(head *ListNode) int {
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    return n
}
// func sortedListToBST(head *ListNode) *TreeNode {
//     var dfs func(curr *ListNode) *TreeNode
//     dfs = func(curr *ListNode) *TreeNode {
//         // base
//         if curr == nil {return nil}
//         if curr.Next == nil {
//             return &TreeNode{Val: curr.Val}
//         }

//         // logic
//         var prev *ListNode
//         slow := curr
//         fast := curr
//         for fast != nil && fast.Next != nil {
//             prev = slow
//             slow = slow.Next
//             fast = fast.Next.Next
//         }
//         prev.Next = nil
//         root := &TreeNode{Val: slow.Val}
//         root.Left = dfs(curr)
//         root.Right = dfs(slow.Next)
//         return root
//     }
//     return dfs(head)
// }