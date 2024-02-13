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
    globalPath := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        globalPath = (globalPath * 10) + r.Val
        dfs(r.Left)        
        dfs(r.Right)
        if r.Left == nil && r.Right == nil {
            total += globalPath
        }
        globalPath /= 10
    }
    dfs(root)
    return total
}