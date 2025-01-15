/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}        
        // logic
        rt := dfs(r.Right)
        lt := dfs(r.Left)
        rh := r.Right
        lh := r.Left
        if lh != nil {
            r.Right = lh
            r.Left = nil
            lt.Right = rh
        }
        if rt != nil {return rt}
        if lt != nil {return lt}
        return r
        
    }
    dfs(root)
}