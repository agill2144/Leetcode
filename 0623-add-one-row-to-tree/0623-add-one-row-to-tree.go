/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        newRoot := &TreeNode{Val: val}
        newRoot.Left = root
        return newRoot
    }
    
    var dfs func(r *TreeNode, currDepth int)
    dfs = func(r *TreeNode, currDepth int)  {
        // base
        if r == nil {return}
        
        // logic
        if currDepth == depth-1 {
            currLeft := r.Left
            currRight := r.Right
            newLeft := &TreeNode{Val: val}
            newRight := &TreeNode{Val: val}
            r.Left = newLeft
            r.Right = newRight
            newLeft.Left = currLeft
            newRight.Right = currRight
        }
        
        dfs(r.Left, currDepth+1)
        dfs(r.Right, currDepth+1)
                
    }
    dfs(root, 1)
    return root
}