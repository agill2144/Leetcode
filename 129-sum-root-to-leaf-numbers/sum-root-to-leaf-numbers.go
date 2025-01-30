/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    total := 0
    var dfs func(r *TreeNode, path int)
    dfs = func(r *TreeNode, path int) {
        // base
        if r == nil {return}

        // logic
        path = (path * 10) + r.Val
        if r.Left == nil && r.Right == nil {
            total += path
            return
        }
        dfs(r.Left, path)
        dfs(r.Right, path)
    }
    dfs(root, 0)
    return total
}