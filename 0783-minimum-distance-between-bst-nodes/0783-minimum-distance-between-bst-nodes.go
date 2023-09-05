/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func minDiffInBST(root *TreeNode) int {
    minDiff := math.MaxInt64
    var prev *TreeNode
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        if prev != nil {
            if r.Val - prev.Val < minDiff {minDiff = r.Val-prev.Val}
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    return minDiff
}