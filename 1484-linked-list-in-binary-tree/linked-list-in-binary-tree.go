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
func isSubPath(head *ListNode, root *TreeNode) bool {
    if head == nil && root == nil {return true}
    if head == nil || root == nil {return false}
    var dfs func(curr *ListNode, r *TreeNode) bool
    dfs = func(curr *ListNode, r *TreeNode) bool {
        // base
        if curr == nil {return true}
        if r == nil {return false}

        // logic
        // ll matches r
        if curr.Val == r.Val {
            return dfs(curr.Next, r.Left) || dfs(curr.Next, r.Right)
        }
        return false
    }
    return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}