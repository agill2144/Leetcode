/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    sum := 0
    var dfs func(r *TreeNode, n int)
    dfs = func(r *TreeNode, n int) {
        // base
        if r == nil {return}
        
        // logic
        n = n*10+r.Val
        if r.Left == nil && r.Right == nil {
            sum += n
            return
        }
        dfs(r.Left, n)
        dfs(r.Right, n)
    }
    dfs(root, 0)
    return sum
}