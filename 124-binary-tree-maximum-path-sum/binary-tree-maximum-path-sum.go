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
        maxSum = max(maxSum, max(r.Val, max(r.Val+left, max(r.Val+right, r.Val+left+right))))
        return max(r.Val, max(left+r.Val, right+r.Val))
    }
    dfs(root)
    return maxSum
}