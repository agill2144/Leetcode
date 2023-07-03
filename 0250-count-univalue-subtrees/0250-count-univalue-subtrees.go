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
        
        // "A uni-value subtree means all nodes of the subtree have the same value."
        // if a node has both childs, they must ALL be the same value
        if r.Left != nil && r.Right != nil && left && right {
            if r.Val == r.Left.Val && r.Val == r.Right.Val {
                count++
                return true
            }
        }
        
        if left && r.Left != nil && r.Left.Val == r.Val && r.Right == nil {
            count++
            return true
        }
        if right && r.Right != nil && r.Right.Val == r.Val && r.Left == nil {
            count++
            return true
        }
        
        return false
    }
    dfs(root)
    return count
}