/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var dfs func(r1, r2 *TreeNode) bool
    dfs = func(r1,r2 *TreeNode) bool {
        if r1 == nil && r2 == nil {return true}
        if r1 == nil || r2 == nil {return false}
        return r1.Val == r2.Val && dfs(r1.Left, r2.Right) && dfs(r1.Right, r2.Left) 
    }
    return dfs(root, root)
}