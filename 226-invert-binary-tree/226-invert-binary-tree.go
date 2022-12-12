/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        r.Left, r.Right = r.Right, r.Left
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return root
}