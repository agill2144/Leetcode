/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: level order using dfs
*/
func findBottomLeftValue(root *TreeNode) int {
    maxLevel := 0
    out := -1
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Right, level+1)
        dfs(r.Left, level+1)
        if r.Left == nil && r.Right == nil {
            if level >= maxLevel {
                maxLevel = level
                out = r.Val
            }
        }
    }
    dfs(root, 0)
    return out
}