/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var (
        xParent *TreeNode
        xDepth int
        yParent *TreeNode
        yDepth int
    )
    var dfs func(r, parent *TreeNode, depth int)
    dfs = func(r, parent *TreeNode, depth int) {
        // base
        if r == nil {return}

        // logic
        if r.Val == x {
            xDepth = depth
            xParent = parent
        }
        if r.Val == y {
            yDepth = depth
            yParent = parent
        }

        dfs(r.Left, r, depth+1)
        dfs(r.Right, r, depth+1)
    }
    dfs(root, nil, 0)
    return xParent != yParent && xDepth == yDepth
}