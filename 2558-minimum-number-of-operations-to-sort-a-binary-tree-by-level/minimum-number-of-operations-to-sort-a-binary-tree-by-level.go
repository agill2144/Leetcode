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
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	idxs := map[int]int{}
	for i := 0; i < len(nums); i++ {
		idxs[nums[i]] = i
	}
	swaps := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == sorted[i] {
			continue
		}
        currVal := nums[i]
        needVal := sorted[i]
        needValOldIdx := idxs[needVal]
        nums[i], nums[needValOldIdx] = nums[needValOldIdx], nums[i]
        idxs[currVal] = needValOldIdx
		swaps++
	}
	return swaps
}