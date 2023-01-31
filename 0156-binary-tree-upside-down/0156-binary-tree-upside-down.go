/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    var newRoot *TreeNode
    var dfs func(r *TreeNode) 
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        if newRoot == nil {
            newRoot = r
        }
        if r.Left != nil  {
            left := r.Left
            right := r.Right
            left.Left = right
            left.Right = r
            r.Left = nil
            r.Right = nil
        }
        dfs(r.Right)
    }
    dfs(root)
    return newRoot
}