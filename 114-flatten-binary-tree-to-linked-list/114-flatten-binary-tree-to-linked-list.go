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
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode){
        // base
        if r == nil {return}
        // logic
        dfs(r.Left)
        if r.Left != nil {
            tmpRight := r.Right
            r.Right = r.Left
            r.Left = nil
            for r.Right != nil {
                r = r.Right
            }
            r.Right = tmpRight
        }
        dfs(r.Right)
    }
    dfs(root)
}