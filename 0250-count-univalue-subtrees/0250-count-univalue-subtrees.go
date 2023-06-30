/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        if r.Left == nil && r.Right == nil {
            count++
            return true
        }
        if r.Left != nil && r.Right != nil && left && right {
            if r.Val == r.Left.Val && r.Val == r.Right.Val {
                count++
                return true
            }
        }
        if r.Left == nil && r.Right != nil && right {
            if r.Val == r.Right.Val {
                count++
                return true
            }
        }
        
        if r.Left != nil && r.Right == nil && left {
            if r.Val == r.Left.Val {
                count++
                return true
            }
        }
        
        return false
    }
    dfs(root)
    return count
}