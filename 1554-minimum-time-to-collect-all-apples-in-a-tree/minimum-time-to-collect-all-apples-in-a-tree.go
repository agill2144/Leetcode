func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    total := 0
    var dfs func(curr, prev int) bool
    dfs = func(curr, prev int) bool {
        // base 

        // logic
        yes := false
        for _, nei := range adjList[curr] {
            if nei == prev {continue}
            if dfs(nei, curr) {
                yes = true
                total += 2
            }
        }
        return hasApple[curr] || yes 
    }
    dfs(0, -1)
    return total
}