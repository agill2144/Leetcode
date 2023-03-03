func validPath(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
    }
    visited := map[int]struct{}{}

    var dfs func(node int, prev int) bool
    dfs = func(node int, prev int) bool {
        // base
        if node == destination {return true}
        if _, ok := visited[node]; ok {return false }
        
        // logic
        visited[node] = struct{}{}
        for _, neighbor := range adjList[node] {
            if neighbor == prev {continue}
            found := dfs(neighbor, node)
            if found {return true}
        }
        return false
    }
    return dfs(source, -1)
}
