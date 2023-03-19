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
        if r == nil {return false}
        
        // logic
        left := dfs(r.Left)
        if !left {r.Left = nil}

        right := dfs(r.Right)
        if !right {r.Right = nil}
                
        return r.Val == 1 || left || right 
    }
    if ok := dfs(root); ok {return root}
    return nil
}