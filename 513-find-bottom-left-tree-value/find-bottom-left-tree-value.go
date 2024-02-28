/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    if root == nil {return -1}
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q)
        leftMost := math.MinInt64
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if leftMost == math.MinInt64 {
                leftMost = dq.Val
            }
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        if len(q) == 0 {return leftMost}
    }
    return -1
}