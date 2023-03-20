/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}
        
        
        // logic
        left := dfs(r.Left)
        if left {r.Left = nil}
        right := dfs(r.Right)
        if right {r.Right = nil}
        if left && right && r.Val == 0 {return true}
        return false
    }
    ok := dfs(root)
    if ok && root.Val == 0 {return nil}
    return root
}