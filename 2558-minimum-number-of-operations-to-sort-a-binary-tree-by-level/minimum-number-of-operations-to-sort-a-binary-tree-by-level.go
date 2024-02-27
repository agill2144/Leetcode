/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	result := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		var lvlNums []int
		for _, node := range q {
			lvlNums = append(lvlNums, node.Val)
		}

		sortedNums := make([]int, len(lvlNums))
		copy(sortedNums, lvlNums)

		sort.Ints(sortedNums)
		for idx, val := range lvlNums {
			lvlNums[idx] = sort.SearchInts(sortedNums, val)
		}

		visited := make([]bool, len(lvlNums))
		for _, num := range lvlNums {
			if visited[num] {
				continue
			}

			tmp := -1
			for false == visited[num] {
				visited[num] = true
				tmp++
				num = lvlNums[num]
			}

			result += tmp
		}

		tmp := q
		q = nil
		for _, node := range tmp {
			if nil != node.Left {
				q = append(q, node.Left)
			}

			if nil != node.Right {
				q = append(q, node.Right)
			}
		}
	}

	return result
}