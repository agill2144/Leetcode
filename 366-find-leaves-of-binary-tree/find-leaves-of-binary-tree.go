/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        maxH := max(left, right)
        if len(out) == maxH {out = append(out, []int{})}
        out[maxH] = append(out[maxH], r.Val)
        return maxH+1
    }
    dfs(root)
    return out
}