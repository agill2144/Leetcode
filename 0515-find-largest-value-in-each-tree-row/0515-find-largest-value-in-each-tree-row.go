/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil {return nil}
    out := []int{}
    q := []*TreeNode{root}
    for len(q) != 0 {
        levelMax := math.MinInt64
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.Val > levelMax {levelMax = dq.Val}
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        out = append(out, levelMax)
    }
    return out
}