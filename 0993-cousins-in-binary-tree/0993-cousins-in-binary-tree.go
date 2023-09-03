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
        xLevel int
        yParent *TreeNode
        yLevel int
    )
    
    var dfs func(r *TreeNode, p *TreeNode, level int)
    dfs = func(r *TreeNode, p *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        if r.Val == x {
            xParent = p
            xLevel = level
        }
        if r.Val == y {
            yParent = p
            yLevel = level
        }
        
        if xParent != nil && yParent != nil {return}
        
        dfs(r.Left, r, level+1)
        dfs(r.Right, r, level+1)
    }
    dfs(root, nil, 0)
    
    return xParent != nil && yParent != nil && xLevel == yLevel && xParent != yParent
    
}