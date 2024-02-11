/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        dfs(r.Left)
        dfs(r.Right)

        
        if r.Left != nil {
            left := r.Left
            right := r.Right
            r.Left = nil
            r.Right = left
            tmp := r
            for tmp.Right != nil {
                tmp = tmp.Right
            }
            tmp.Right = right
        }
        
    }
    
    dfs(root)
}