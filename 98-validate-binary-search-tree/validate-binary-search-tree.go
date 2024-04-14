/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}

        // logic
        lok := dfs(r.Left)
        if !lok {return false}
        if prev != nil && prev.Val >= r.Val {return false}
        prev = r
        return dfs(r.Right)
    }
    return dfs(root)
}