func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
    }
    visited := make([]bool, n)
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        // leaf node thats not destination, return false
        if len(adjList[node]) == 0 && node != destination {return false}
        if visited[node] {return true}
        
        
        // logic
        visited[node] = true
        path[node] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {
                return false
            }
        }
        path[node] = false
        return true
    }
    return dfs(source, make([]bool, n))
}