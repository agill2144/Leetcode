/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var paths []int
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        paths = append(paths, r.Val)
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return paths
}