func validTree(n int, edges [][]int) bool {
    // valid tree = all components are connected, no cycles
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }

    p := make([]bool, n)
    visited := make([]bool, n)
    visitedCount := 0
    var dfs func(node, prev int, path []bool) bool
    dfs = func(node, prev int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}

        // logic
        visited[node] = true
        path[node] = true
        visitedCount++
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if !dfs(nei, node, path) {return false}
        }
        path[node] = false
        return true        
    }

    if !dfs(0, -1, p) {return false}
    return visitedCount == n
}