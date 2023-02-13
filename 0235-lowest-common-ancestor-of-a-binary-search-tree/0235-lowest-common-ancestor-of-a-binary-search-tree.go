/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
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
        if (p.Val > r.Val && q.Val < r.Val) || (p.Val < r.Val && q.Val > r.Val) {return r}
        if p.Val > r.Val && q.Val > r.Val {
            if node := dfs(r.Right); node != nil {return node}
        }
        return dfs(r.Left)
    }
    return dfs(root)
}