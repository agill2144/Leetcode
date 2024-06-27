/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}        
        // logic
        if r.Val >= low && r.Val <= high {
            sum += r.Val
        }
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return sum
}