func minimumOperations(root *TreeNode) int {
	ops := 0
	for currNodes := []*TreeNode{root}; len(currNodes) > 0; {
		nextNodes := make([]*TreeNode, 0, len(currNodes))
		nums := make([]int, 0, len(currNodes))
		for _, v := range currNodes {
			nums = append(nums, v.Val)
			if v.Left != nil {
				nextNodes = append(nextNodes, v.Left)
			}
			if v.Right != nil {
				nextNodes = append(nextNodes, v.Right)
			}
		}
		ops += op(nums)
		currNodes = nextNodes
	}
	return ops
}

func op(nums []int) int {
	swap := 0
	posMap := make(map[int]int, len(nums))
	for i, v := range nums {
		posMap[v] = i
	}
	numsSorted := make([]int, len(nums))
	copy(numsSorted, nums)
	sort.Ints(numsSorted)
	for i := range numsSorted {
		if nums[i] != numsSorted[i] {
			// swap once
			swap++
			wrongVal := nums[i]
			wrongPos := posMap[numsSorted[i]]
			nums[i], nums[wrongPos] = nums[wrongPos], nums[i]
			posMap[nums[i]] = i
			posMap[wrongVal] = wrongPos
		}
	}
	return swap
}