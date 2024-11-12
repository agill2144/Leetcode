/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    out := []int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        // logic
        if len(out) == level {
            out = append(out, r.Val)
        }
        dfs(r.Right, level+1)
        dfs(r.Left, level+1)
    }
    dfs(root, 0)
    return out
}