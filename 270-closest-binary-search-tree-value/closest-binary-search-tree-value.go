/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    closest := root.Val
    closestDiff := abs(float64(root.Val)-target)
    for root != nil {
        diff := abs(float64(root.Val)-target)
        if diff < closestDiff {closest = root.Val; closestDiff = diff}
        if diff == closestDiff {closest = min(root.Val, closest)}
        
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return closest
}

func abs(x float64) float64 {
    if x < 0 {return x * -1.0}
    return x
}