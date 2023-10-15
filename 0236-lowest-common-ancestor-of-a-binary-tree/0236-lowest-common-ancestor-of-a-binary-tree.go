/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}
        
        // logic
        left := dfs(r.Left)
        if r == p || r == q {return r}

        right := dfs(r.Right)
        
        if left == nil && right == nil {return nil}
        if left != nil && right != nil {return r}
        if left != nil && right == nil {return left}
        return right
    }
    return dfs(root)
}