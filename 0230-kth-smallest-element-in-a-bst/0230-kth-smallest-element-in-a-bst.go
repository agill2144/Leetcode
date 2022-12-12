/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var dfs func (r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return -1}        
        // logic
        left := dfs(r.Left)
        if left >= 0 {
            return left
        }
        k--
        if k == 0 {
            return r.Val
        }
        return dfs(r.Right)
    }
    return dfs(root)
}