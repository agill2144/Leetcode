/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    if root == nil {return 0}
    var res int = math.MaxInt64
    var resDiff float64 = math.MaxFloat64
    for root != nil {
        diff := abs(float64(root.Val)-target)
        if diff == resDiff { if root.Val < res {res = root.Val} }
        if diff < resDiff {res = root.Val; resDiff = diff}
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return res
}

func abs(x float64) float64 {
    if x < 0.0 {return x*-1.0}
    return x
}