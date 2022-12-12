/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {return 0}
    maxDiameter := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        
        // base
        if r == nil {return 0}
        
        // logic
        leftDepth := dfs(r.Left)
        rightDepth := dfs(r.Right)
        
        maxDiameter = max(maxDiameter, leftDepth+rightDepth)
        
        return max(leftDepth, rightDepth)+1
        
    }
    dfs(root)
    return maxDiameter
    
}

func max(x, y int) int {
    if x > y {return x}
    return y
}