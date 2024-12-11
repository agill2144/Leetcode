/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}
    n := 0
    var count func(r *TreeNode)
    count = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        n++
        count(r.Left)
        count(r.Right)
    }
    count(root)
    var dfs func(r *TreeNode, idx int) bool
    dfs = func(r *TreeNode, idx int) bool {
        // base
        if r == nil {return true}

        // logic
        if idx >= n {return false}
        if !dfs(r.Left, 2*idx+1) {return false}
        return dfs(r.Right, 2*idx+2)
    }
    return dfs(root,0)
}