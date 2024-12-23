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
        ogIdx := idxs[needVal]
        nums[i], nums[ogIdx] = nums[ogIdx], nums[i]
        // we swapped curr val with correct val
        // and moved current value to somewhere to the right
        // therefore update our state / idx map that the new idx of current value is ogIdx
        idxs[currVal] = ogIdx
		swaps++
	}
	return swaps
}