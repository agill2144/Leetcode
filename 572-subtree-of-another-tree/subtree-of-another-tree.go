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
            if isSame(r, subRoot) {return true}
        }
        if dfs(r.Left) {return true}
        if dfs(r.Right) {return true}
        return false
    }
    return dfs(root)
}

func isSame(a, b *TreeNode) bool {
    if a == nil && b == nil {return true}
    if a == nil || b == nil {return false}
    if a.Val != b.Val {return false}
    return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}