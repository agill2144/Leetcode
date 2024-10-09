/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {return}
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        if left != nil {
            left.Right = r.Right
            r.Right = r.Left
            r.Left = nil
        }
        if right != nil {return right}
        if left != nil {return left}
        return r
    }
    dfs(root)
}