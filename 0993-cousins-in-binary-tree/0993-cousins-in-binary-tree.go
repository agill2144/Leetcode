/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var xParent *TreeNode
    xDepth := -1
    var yParent *TreeNode
    yDepth := -1
    var dfs func(r *TreeNode, parent *TreeNode, depth int) 
    dfs = func(r *TreeNode, parent *TreeNode, depth int) {
        // base
        if r == nil {return}
        
        // logic
        if xParent != nil && yParent != nil {return}
        
        if r.Val == x {
            xParent = parent
            xDepth = depth
        }
        if r.Val == y {
            yParent = parent
            yDepth = depth
        }
        
        dfs(r.Left, r, depth+1)
        dfs(r.Right, r, depth+1)
    }
    
    dfs(root, nil, 0)
    return (xDepth != -1 && yDepth != -1) && (xDepth == yDepth && xParent != yParent)
    
}