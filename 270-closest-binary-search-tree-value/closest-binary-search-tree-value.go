/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    minVal := math.MaxInt64
    var minDiff float64 = math.MaxInt64
    for root != nil {
        diff := abs(float64(root.Val)-target)
        if diff == minDiff {
            if root.Val < minVal {minVal = root.Val}
        }
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

func abs(x float64) float64 {
    if x < 0.0 {return x*-1.0}
    return x
}