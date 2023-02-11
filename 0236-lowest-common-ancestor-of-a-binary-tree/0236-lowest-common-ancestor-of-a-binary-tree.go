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
        if r == p || r == q {return r}
        left := dfs(r.Left)
        right := dfs(r.Right)

        if left != nil && right != nil {
            return r
        }

        if left == nil && right == nil {return nil}
        if left != nil && right == nil {return left}
        if left == nil && right != nil {return right}
        return nil
    }
    return dfs(root)
}