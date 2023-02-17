/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    diff := math.MaxInt64
    var dfs func(r *TreeNode) 
    var prev *TreeNode
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        dfs(r.Left)
        if prev != nil {
            if r.Val - prev.Val < diff {diff = r.Val-prev.Val}
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    return diff
}