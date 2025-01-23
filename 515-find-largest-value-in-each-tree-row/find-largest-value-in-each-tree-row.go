/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    out := []int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        if len(out) == level {
            out = append(out, r.Val)
        } else if r.Val > out[level] {
            out[level] = r.Val
        }
        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    dfs(root, 0)
    return out
}