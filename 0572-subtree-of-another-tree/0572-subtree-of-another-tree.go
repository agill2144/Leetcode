/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return false}
        
        // logic
        if r.Val == subRoot.Val {
            if isSameTree(r, subRoot) {return true}
        }
        if ok := dfs(r.Left); ok {return true}
        return dfs(r.Right)
    }
    return dfs(root)
}

func isSameTree(a, b *TreeNode) bool {
    if a == nil && b == nil {return true}
    if (a == nil && b != nil) || (a != nil && b == nil) {return false}
    return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}