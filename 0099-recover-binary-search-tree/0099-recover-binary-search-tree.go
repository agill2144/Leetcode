/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    if root == nil {return}
    var first *TreeNode
    var second *TreeNode
    var prev *TreeNode
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        
        // logic
        dfs(r.Left)
        if prev != nil {
            if prev.Val > r.Val {
                if first == nil {
                    first = prev
                }
                second = r
                
            }
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    first.Val, second.Val = second.Val, first.Val

}