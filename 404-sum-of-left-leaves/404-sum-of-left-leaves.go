/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
            sum += r.Left.Val
        }
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return sum
}