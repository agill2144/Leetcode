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
        leftRightRoot := left + right + r.Val
        rootLeft := r.Val + left
        rootRight := r.Val + right
        maxSum = max(maxSum,max(leftRightRoot,max(rootLeft,max(rootRight, r.Val))))
        return max(rootLeft, max(r.Val, rootRight))
    }
    dfs(root)
    return maxSum
}