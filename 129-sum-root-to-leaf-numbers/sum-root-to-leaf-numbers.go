/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    total := 0
    var dfs func(r *TreeNode, p int)
    dfs = func(r *TreeNode, p int) {
        // base
        if r == nil {return}
        
        // logic
        p = p * 10 + r.Val
        dfs(r.Left, p)
        dfs(r.Right, p)
        if r.Left == nil && r.Right == nil {
            total += p
        }   
    }
    dfs(root, 0)
    return total
}