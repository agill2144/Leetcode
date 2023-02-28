func findCircleNum(isConnected [][]int) int {
    adjList := map[int][]int{}
    n := len(isConnected)
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && i != j {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }
    visited := map[int]struct{}{}
    var dfs func(node int) 
    dfs = func(node int) {
        // base
        if _, ok := visited[node]; ok { return }
        
        // logic
        visited[node] = struct{}{}
        for _, edge := range adjList[node] {
            if _, ok := visited[edge]; !ok {dfs(edge)}
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            count++
            dfs(i)
        }
    }
    return count
}
