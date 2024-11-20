/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {return true}
    var dfs func(r *TreeNode, minVal, maxVal int) bool
    dfs = func(r *TreeNode, minVal, maxVal int) bool {
        // base
        if r == nil {return true}

        // logic
        if r.Val <= minVal || r.Val >= maxVal {return false}
        if !dfs(r.Left, minVal, r.Val) {return false}
        if !dfs(r.Right, r.Val, maxVal) {return false}
        return true
    }
    return dfs(root, math.MinInt64, math.MaxInt64)
}