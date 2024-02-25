/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    minDiff := math.MaxFloat64
    minVal := math.MaxInt64

    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        diff := abs(float64(r.Val) - target)

        if diff < minDiff || diff == minDiff {
            if diff < minDiff { minDiff = diff; minVal = r.Val}
            if diff == minDiff { minVal = min(r.Val, minVal) }
        }

        if target > float64(r.Val) {
            dfs(r.Right)
        } else {
            dfs(r.Left)
        }
    }
    dfs(root)
    return minVal
}

func abs(x float64) float64 {
    if x < 0.0 {return x * -1.0}
    return x
}