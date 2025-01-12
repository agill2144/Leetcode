func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        a, b := edges[i][0], edges[i][1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)
    }
    var dfs func(node, prev int) int
    dfs = func(node, prev int) int {
        // base

        // logic
        total := 0
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            c := dfs(nei, node)
            total += c
            if c > 0 || hasApple[nei] {total += 2}
        }
        return total
    }
    return dfs(0,-1)
}