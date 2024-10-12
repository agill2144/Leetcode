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
        if r == nil {return nil}

        // logic
        lt := dfs(r.Left)
        rt := dfs(r.Right)
        lh := r.Left
        rh := r.Right
        if lh != nil {
            r.Right = lh
            lt.Right = rh
            r.Left = nil
        }
        if rt != nil {return rt}
        if lt != nil {return lt}
        return r
    }
    dfs(root)
}