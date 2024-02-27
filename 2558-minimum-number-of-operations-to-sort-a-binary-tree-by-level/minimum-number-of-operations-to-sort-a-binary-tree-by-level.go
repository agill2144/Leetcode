func minimumOperations(root *TreeNode) int {
	result, queue := 0, []*TreeNode{root}
	for len(queue) > 0 {
		size, nums := len(queue), make([]int, 0)
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:]
			nums = append(nums, t.Val)
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		result += cnt(nums)
	}
	return result
}

func cnt(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	index := make(map[int]int)
	for i, num := range nums {
		index[num] = i
	}
	result := 0
	for i, num := range nums {
		if num != sorted[i] {
			nums[i], nums[index[sorted[i]]] = nums[index[sorted[i]]], nums[i]
			index[nums[i]], index[num] = i, index[sorted[i]]
			result++
		}
	}
	return result
}