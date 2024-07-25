/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    total := 0
    var dfs func(r, p, g *TreeNode)
    dfs = func(r, p, g *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if g != nil && g.Val % 2 == 0 {
            total += r.Val
        }
        dfs(r.Left, r, p)
        dfs(r.Right, r, p)
    }
    dfs(root, nil, nil)
    return total
}