/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
    totalSwaps := 0
    if root == nil {
        return totalSwaps
    }
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q)
        level := []int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            level = append(level, dq.Val)
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        totalSwaps += countSwaps(level)
    }
    return totalSwaps
}

func countSwaps(nums []int) int {
    n := len(nums)
    valIdx := [][]int{}
    for i := 0; i < n; i++ {
        valIdx = append(valIdx, []int{nums[i], i})
    }
    sort.Slice(valIdx, func(i, j int)bool{
        return valIdx[i][0] < valIdx[j][0]
    })
    swaps := 0
    for i := 0; i < n; i++ {
        idx := valIdx[i][1]
        for idx != i {
            valIdx[i], valIdx[idx] = valIdx[idx], valIdx[i]
            idx = valIdx[i][1]
            swaps++
        }
    }
    return swaps
}