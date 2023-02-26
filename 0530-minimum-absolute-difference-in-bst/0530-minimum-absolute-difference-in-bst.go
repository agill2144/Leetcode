/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    var prev *TreeNode
    min := math.MaxInt64
    var dfs func(r *TreeNode) 
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}        
        // logic
        dfs(r.Left)
        if prev != nil {
            if r.Val - prev.Val < min {
                min = r.Val - prev.Val
            }
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    return min
}