/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    xLevel, yLevel := -1, -1
    var xParent *TreeNode
    var yParent *TreeNode
    var dfs func(r, p *TreeNode, level int)
    dfs = func(r, p *TreeNode, level int) {
        // base
        if r == nil {return}


        // logic
        if r.Val == x {xLevel = level; xParent = p}
        if r.Val == y {yLevel = level; yParent = p}
        if xParent != nil && yParent != nil {return}
        dfs(r.Left, r, level+1)
        dfs(r.Right, r, level+1)
    }
    dfs(root, root, 0)
    if xParent == nil || yParent == nil {return false}
    return xLevel == yLevel && xParent != yParent
}