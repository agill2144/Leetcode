import (
	"container/list"
	"sort"
)

func minimumOperations(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	result := 0

	for {
		n := queue.Len()
		if n == 0 {
			break
		}
		arr := make([]int, 0, n)
		for i := 0; i < n; i++ {
			elem := queue.Front()
            queue.Remove(elem)
            e := elem.Value.(*TreeNode)

			arr = append(arr, e.Val)
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
        qty := minPermutations(arr)
        result += qty
	}

	return result
}

func minPermutations(arr []int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	idxMap := make(map[int]int, len(arr))
	permutationsQty := 0

	for i := 0; i < len(arr); i++ {
		idxMap[arr[i]] = i
	}

	for i, elem := range arr {
		if elem != sorted[i] {
			index := idxMap[sorted[i]]
			arr[i], arr[index] = arr[index], arr[i]
			idxMap[arr[i]] = i
			idxMap[arr[index]] = index
			permutationsQty++
		}
	}
	return permutationsQty
}