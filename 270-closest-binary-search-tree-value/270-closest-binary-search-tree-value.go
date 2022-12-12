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
    
    for root != nil {
        diff := abs(float64(root.Val) - target)
        if diff < minDiff {
            minDiff = diff
            minVal = root.Val
        }
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return minVal
}

func abs(n float64) float64  {
    if n < 0 {return n * -1.0}
    return n
} 