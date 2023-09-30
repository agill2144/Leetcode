/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// bottom up
func equalToDescendants(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        total := left+right
        if total == r.Val {count++}
        return total + r.Val
    }
    dfs(root)
    return count
}