func findCircleNum(isConnected [][]int) int {
    m := len(isConnected)
    n := len(isConnected[0])
    adjList := map[int][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && i != j {
                adjList[i] = append(adjList[i], j)
            }
        }
    }
    visited := make([]bool, m)
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            dfs(nei)
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
            count++
        }
    }
    return count
}