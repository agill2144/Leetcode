/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return true}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)

        // leaf node
        if r.Left == nil && r.Right == nil {
            if r.Val == 1 {return true}
            return false
        }

        // some middle node
        if r.Val == 2 {
            return left || right
        }
        return left && right
    }
    return dfs(root)
}