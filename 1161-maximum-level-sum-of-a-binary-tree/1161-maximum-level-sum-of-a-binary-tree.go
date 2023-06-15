/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    ansLevel := 0
    maxSum := math.MinInt64
    q := []*TreeNode{root}
    level := 1
    for len(q) != 0 {
        sum := 0
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            sum += dq.Val
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        if sum > maxSum {
            maxSum = sum
            ansLevel = level
        } 
        level++
    }
    return ansLevel
}