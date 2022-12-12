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
        if r == nil {
            return 0
        }
        // logic
        // why max with 0?
        // because we may have a subtree bottom up whose ans is negative and not choosing that subtree all together is better
        // therefore max between left, 0
        left := max(dfs(r.Left),0)
        right := max(dfs(r.Right),0)
        
        totalSum := r.Val + left + right
        maxSum = max(maxSum, totalSum)
        return max(r.Val+left, r.Val+right)
    }
    dfs(root)
    return maxSum
}

func max(x, y int) int{
    if x > y {return x}
    return y
}