/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maximumAverageSubtree(root *TreeNode) float64 {
    max := 0.0
    var dfs func(r *TreeNode) (int, int)
    dfs = func(r *TreeNode) (int, int) {
        // base
        if r == nil {return 0,0}
        
        // logic
        leftSum, leftCount := dfs(r.Left)
        rightSum, rightCount := dfs(r.Right)
        totalNodes := leftCount+rightCount+1
        sum := leftSum+rightSum+r.Val
        
        avg := float64(sum) / float64(totalNodes)
        if avg > max {
            max = avg
        }
        
        return sum, totalNodes
    }
    _, _ = dfs(root)
    return max
}
