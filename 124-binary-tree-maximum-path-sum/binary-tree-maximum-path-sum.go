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
        total := r.Val + left + right
        leftTotal := r.Val + left
        rightTotal := r.Val + right
        maxSum = max(leftTotal, max(rightTotal, max(total, max(r.Val, maxSum))))
        return max(leftTotal, max(rightTotal, r.Val))
    }
    dfs(root)
    return maxSum
}