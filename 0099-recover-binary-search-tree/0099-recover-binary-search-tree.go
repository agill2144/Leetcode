/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    var prev *TreeNode
    var first *TreeNode
    var second *TreeNode
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
    if first != nil {
        first.Val , second.Val = second.Val, first.Val
    }

}