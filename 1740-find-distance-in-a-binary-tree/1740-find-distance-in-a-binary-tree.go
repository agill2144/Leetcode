/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDistance(root *TreeNode, p int, q int) int {
    var lca func(r *TreeNode) *TreeNode
    lca = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}
        
        // logic
        if r.Val == p || r.Val == q {
            return r
        }
        left := lca(r.Left)
        right := lca(r.Right)
        
        if left != nil && right != nil {return r}
        if left == nil && right != nil {return right}
        if left != nil && right == nil {return left}
        return nil
    }
    lowestCommon := lca(root)
    
    totalDist := 0
    var dfs func(r *TreeNode, target int, level int) bool
    dfs = func(r *TreeNode, target int, level int) bool {
        // base
        if r == nil {return false}
        
        // logic
        if r.Val == target {
            totalDist += level
            return true
        }
        if ok := dfs(r.Left, target, level+1); ok {return true}
        if ok := dfs(r.Right, target, level+1); ok {return true}
        return false
    }
    dfs(lowestCommon, p, 0)
    dfs(lowestCommon, q, 0)
    return totalDist
}
