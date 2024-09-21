/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt64
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        total := left + right + r.Val
        maxSum = max(maxSum, max(r.Val+left, max(r.Val+right, max(r.Val, total))))
        return max(r.Val, max(r.Val+left, r.Val+right))
    }
    dfs(root)
    return maxSum
}