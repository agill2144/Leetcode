/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    out := 0
    var dfs func(r *TreeNode, h int)
    dfs = func(r *TreeNode, h int) {
        // base
        if r == nil {return}

        // logic
        h++
        if r.Left == nil && r.Right == nil {
            out = max(h, out)
        }
        dfs(r.Left, h)
        dfs(r.Right, h)
    }
    dfs(root, 0)
    return out
}