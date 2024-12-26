/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var dfs func(r *TreeNode) *TreeNode
    dfs = func(r *TreeNode) *TreeNode{
        // base
        if r == nil {return nil}

        // logic
        left := dfs(r.Left)
        if left != nil {return left}
        k--
        if k == 0 {return r}
        return dfs(r.Right)
    }
    return dfs(root).Val
}