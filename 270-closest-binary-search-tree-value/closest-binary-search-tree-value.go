/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    var minDiff float64 = math.MaxInt64
    res := -1
    for root != nil {
        diff := abs(float64(root.Val)-target)
        if diff < minDiff {
            minDiff = diff
            res = root.Val
        } else if diff == minDiff && root.Val < res {
            res = root.Val
        }
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return res
}

func abs(x float64) float64 {
    if x < 0 {return x*-1.0}
    return x
}