/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var dfs func(r *TreeNode, min, max int) bool
    dfs = func(r *TreeNode, min, max int) bool {
        // base
        if r == nil {return true}

        // logic
        if r.Val <= min || r.Val >= max {return false}
        if !dfs(r.Left, min, r.Val) {return false}
        return dfs(r.Right, r.Val, max)
    }
    return dfs(root, math.MinInt64, math.MaxInt64)
}