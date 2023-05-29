/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    var dfs func(r *TreeNode) 
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Right)
        sum += r.Val
        r.Val = sum
        dfs(r.Left)
    }
    dfs(root)
    return root
}