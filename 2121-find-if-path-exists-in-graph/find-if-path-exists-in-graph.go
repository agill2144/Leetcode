func validPath(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    visited := make([]bool, n)
    var dfs func(node, prev int) bool
    dfs = func(node, prev int) bool {
        // base
        if node == destination {return true}
        if visited[node] {return false}

        // logic
        visited[node] = true
        for _, nei := range adjList[node]{
            if nei == prev {continue}
            if dfs(nei, node) {return true}
        }
        return false
    }
    return dfs(source, -1)
}