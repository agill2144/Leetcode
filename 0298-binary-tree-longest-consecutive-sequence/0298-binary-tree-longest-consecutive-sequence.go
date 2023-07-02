/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
    out := 0
    var dfs func(r *TreeNode, size int)
    dfs = func(r *TreeNode, size int) {
        // base
        if size > out {out = size}
        if r == nil {return}
        
        // logic
        if r.Left != nil {
            if  r.Val+1 == r.Left.Val {
                dfs(r.Left, size+1)
            } else {
                dfs(r.Left, 1)
            }
        }
        if r.Right != nil {
            if  r.Val+1 == r.Right.Val {
                dfs(r.Right, size+1)
            } else {
                dfs(r.Right, 1)
            }
        }
    }
    dfs(root, 1)
    return out
}