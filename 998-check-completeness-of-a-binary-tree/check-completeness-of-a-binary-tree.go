/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    n := size(root)
    var dfs func(r *TreeNode, idx int) bool
    dfs = func(r *TreeNode, idx int) bool {
        // base
        if r == nil {return true}

        // logic
        if idx >= n {return false}
        if !dfs(r.Left, 2*idx+1) {return false}
        if !dfs(r.Right, 2*idx+2) {return false}
        return true
    }
    return dfs(root, 0)
}

func size(root *TreeNode) int {
    if root == nil {return 0}
    count := 0 
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return }

        // logic
        count++
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return count
}