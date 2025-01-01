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
    var diff float64 = abs(float64(root.Val)-target)
    res := root.Val
    for root != nil {
        currDiff := abs(float64(root.Val)-target)
        
        if currDiff == diff && root.Val < res {res = root.Val}
        if currDiff < diff {
            diff = currDiff
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
    if x < 0.0 {return x*-1.0}
    return x
}