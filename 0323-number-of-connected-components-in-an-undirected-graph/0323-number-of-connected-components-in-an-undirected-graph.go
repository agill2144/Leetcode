func countComponents(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        node := edges[i][0]
        edge := edges[i][1]
        adjList[node] = append(adjList[node], edge)
        adjList[edge] = append(adjList[edge], node)
    }
    visited := map[int]bool{}
    count := 0
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if _, ok := visited[node]; ok {return}

        // logic
        visited[node] = true
        for _, edge := range adjList[node] {
            dfs(edge)
        }
    }
    
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            dfs(i)
            count++
        }
    }
    return count
}