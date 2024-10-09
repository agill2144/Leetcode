/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {return}
    var prev *TreeNode
    var dfs func(r *TreeNode)   
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Right)
        dfs(r.Left)
        r.Right = prev
        r.Left = nil
        prev = r
    }
    dfs(root)
}