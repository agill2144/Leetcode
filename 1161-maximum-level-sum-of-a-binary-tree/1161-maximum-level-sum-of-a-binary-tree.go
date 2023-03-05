/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    q := []*TreeNode{root}
    levels := 0
    ans := 0
    maxSum := math.MinInt64
    for len(q) != 0 {
        qSize := len(q)
        levelSum := 0
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            levelSum += dq.Val
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        levels++
        if levelSum > maxSum {
            maxSum = levelSum
            ans = levels
        }
    }
    return ans
}