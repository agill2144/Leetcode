/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(r1, r2 *TreeNode) bool
    dfs = func(r1, r2 *TreeNode ) bool {
        if r1 == nil && r2 == nil {return true}
        if r1 == nil || r2 == nil {return false}
        if r1.Val != r2.Val {return false}

        left := dfs(r1.Left, r2.Left)
        if !left {return false}
        right := dfs(r1.Right, r2.Right)
        if !right {return false}
        
        return true
    }
    return dfs(p,q)
}