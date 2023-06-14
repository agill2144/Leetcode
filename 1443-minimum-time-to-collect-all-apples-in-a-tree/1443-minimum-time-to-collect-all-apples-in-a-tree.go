func minTime(n int, edges [][]int, hasApple []bool) int {
	connections := make([][]int32, n)
	for _, edge := range edges {
		connections[edge[0]] = append(connections[edge[0]], int32(edge[1]))
		connections[edge[1]] = append(connections[edge[1]], int32(edge[0]))
	}
	var dfs func(nodeIdx int32) int32
	dfs = func(nodeIdx int32) int32 {
		var result int32
		children := connections[nodeIdx]
		connections[nodeIdx] = nil
		for _, child := range children {
			if connections[child] != nil {
				result += dfs(child)
			}
		}
		if result != 0 || hasApple[nodeIdx] {
			result++
		}
		return result
	}
	result := int(dfs(0))
	if result != 0 {
		result = (result - 1) << 1
	}
	return result
}