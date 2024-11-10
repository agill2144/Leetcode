/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // inorder on a BST = sorted order
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}

        // logic
        left := dfs(r.Left)
        if !left {return false}
        if prev != nil {
            if prev.Val >= r.Val {return false}
        }
        prev = r
        return dfs(r.Right)
    }
    return dfs(root)
}