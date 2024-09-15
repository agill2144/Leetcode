/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// bottom up dfs
func maxPathSum(root *TreeNode) int {
    if root == nil {return 0}
    maxSum := math.MinInt64
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        total := left+right+r.Val
        leftSum := left + r.Val
        rightSum := right + r.Val
        maxSum = max(maxSum, max(total, max(leftSum, max(rightSum, r.Val))))
        return max(leftSum, max(rightSum, r.Val))
    }
    dfs(root)
    return maxSum
}