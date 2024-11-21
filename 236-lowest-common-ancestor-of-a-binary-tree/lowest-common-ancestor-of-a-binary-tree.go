/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p == nil || q == nil || root == nil {return nil}
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}

        // logic
        if r == p || r == q {return r}
        left := dfs(r.Left)
        right := dfs(r.Right)
        if left != nil && right != nil {return r}
        if left != nil {return left}
        return right
    }
    return dfs(root)
}