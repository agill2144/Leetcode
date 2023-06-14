func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1]) 
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0]) 
    }

    count := 0
    var dfs func(node int, prev int) int
    dfs = func(node, prev int) int {
        // base
        // logic
        childHasApples := false
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            childAppleCount := dfs(nei, node)
            if childAppleCount != 0 {
                childHasApples = true
            }
            count += childAppleCount
        }
        if childHasApples || hasApple[node] {return 2}
        return 0
    }
    dfs(0, -1)
    return count
}