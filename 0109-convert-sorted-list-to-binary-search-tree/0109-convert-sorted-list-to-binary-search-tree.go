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
    var dfs func(left, right *ListNode) *TreeNode
    dfs = func(left, right *ListNode) *TreeNode {
        // base
        if left == right {return nil}
        
        // logic
        // find mid given start = left and end = right
        slow := left
        fast := left
        // find mid point
        for fast != right && fast.Next != right {
            slow = slow.Next
            fast = fast.Next.Next
        }
        // mid is at slow now
        root := &TreeNode{Val: slow.Val}
        // left of slow is left subTree
        root.Left = dfs(left, slow)
        // right of slow is right subTree
        root.Right = dfs(slow.Next, right)
        return root
    }
    return dfs(head, nil)
    
}