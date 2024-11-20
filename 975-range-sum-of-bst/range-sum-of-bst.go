/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    total := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if r.Val >= low && r.Val <= high {total += r.Val}
        if r.Val > low {dfs(r.Left)}
        if r.Val < high {dfs(r.Right)}
    }
    dfs(root)
    return total
}