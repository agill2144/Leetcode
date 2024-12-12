/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    minVal := root.Val
    minDiff := abs(float64(root.Val)-target)
    for root != nil {
        diff := abs(float64(root.Val)-target)
        if diff < minDiff {minVal = root.Val; minDiff = diff}
        if diff == minDiff {minVal = min(root.Val, minVal)}
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return minVal
}

func abs(x float64) float64 {
    if x < 0 {return x * -1.0}
    return x
}