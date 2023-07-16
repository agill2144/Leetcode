/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    if root == nil {return 0}
    sum := 0
    var dfs func(r, p, gp *TreeNode)
    dfs = func(r, p, gp *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if gp != nil && gp.Val % 2 == 0 {
            sum += r.Val
        }
        dfs(r.Left, r, p)
        dfs(r.Right, r, p)
    }
    dfs(root, nil, nil)
    
    return sum
}