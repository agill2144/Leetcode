func eventualSafeNodes(graph [][]int) []int {
	safeNode := map[int]bool{}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if safe, f := safeNode[i]; f {
			return safe
		}

		safeNode[i] = false
		for _, nei := range graph[i] {
			if !dfs(nei) {
				return false
			}
		}
		safeNode[i] = true
		return true
	}

	res := []int{}
	for i := 0; i < len(graph); i++ {
		if dfs(i) {
			res = append(res, i)
		}
	}
	return res
}

