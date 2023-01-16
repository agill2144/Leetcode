/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfs func(r *TreeNode, sum int) bool
    dfs = func(r *TreeNode, sum int) bool {
        // base
        if r == nil {return false}
        
        // logic
        sum += r.Val
        left := dfs(r.Left, sum)
        if left {return left}
        
        if sum == targetSum && r.Left == nil && r.Right == nil {return true}
        
        return dfs(r.Right, sum)
    }
    return dfs(root, 0)
}