/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {return 0}
    count := 0
    var dfs func(r *TreeNode, sum int)
    dfs = func(r *TreeNode, sum int) {
        // base
        if r == nil {return}
        
        // logic
        sum += r.Val
        if sum == targetSum {
            count++
        }
        dfs(r.Left, sum)
        dfs(r.Right, sum)
    }
    
    var dfs2 func(r *TreeNode)
    dfs2 = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r, 0)
        dfs2(r.Left)
        dfs2(r.Right)
    }
    
    dfs2(root)
    return count
}

