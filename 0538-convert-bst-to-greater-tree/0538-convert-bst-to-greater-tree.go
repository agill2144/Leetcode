/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    /*
        1. get the total sum
        2. Remove the runningSum when running inorder bst + add curr val
    */
    total := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        total += r.Val
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    
    running := 0
    var dfs2 func(r *TreeNode)
    dfs2 = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs2(r.Left)
        running += r.Val
        r.Val = total-running+r.Val
        dfs2(r.Right)
    }
    dfs2(root)
    return root
    
}